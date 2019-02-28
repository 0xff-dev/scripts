package session

import (
	"container/list"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"
)

// session全局管理
type Manager struct {
	cookieName  string
	lock        sync.RWMutex
	provider    Provider
	maxlifetime int64
}

type Provider interface {
	SessionInit(sid string) (Session, error)
	SessionRead(sid string) (Session, error)
	SessionDestory(sid string) error
	SessionGC(maxlifetime int64)
}

type Session interface {
	Set(key, value interface{}) error
	Get(key interface{}) interface{}
	Delete(key interface{}) error
	SessionID() string
}

var (
	GlobalManager *Manager
	providers     map[string]Provider
	pi            = &ProviderImp{list: list.New()}
)

// Manager
func NewSessionManager(provideName, cookieName string, maxlifetime int64) (*Manager, error) {
	provide, ok := providers[provideName]
	if !ok {
		return nil, fmt.Errorf("unknown provider")
	}
	return &Manager{cookieName: cookieName, provider: provide, maxlifetime: maxlifetime}, nil
}

func (m *Manager) sessionID() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func (m *Manager) SessionStart(writer http.ResponseWriter, req *http.Request) (session Session) {
	m.lock.Lock()
	defer m.lock.Unlock()
	cookie, err := req.Cookie(m.cookieName)
	if err != nil || cookie.Value == "" {
		// 首次来
		sid := m.sessionID()
		session, _ = m.provider.SessionInit(sid)
		cookie := &http.Cookie{Name: m.cookieName, Value: url.QueryEscape(sid),
			Path: "/", HttpOnly: true, MaxAge: int(m.maxlifetime)}
		http.SetCookie(writer, cookie)
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = m.provider.SessionRead(sid)
	}
	return
}

func (m *Manager) SessionDestory(writer http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie(m.cookieName)
	if err != nil || cookie.Value == "" {
		return
	}
	m.lock.Lock()
	defer m.lock.Unlock()
	sid, _ := url.QueryUnescape(cookie.Value)
	m.provider.SessionDestory(sid)
	expiration := time.Now()
	cookie = &http.Cookie{Name: m.cookieName, Path: "/", HttpOnly: true, Expires: expiration, MaxAge: -1}
	http.SetCookie(writer, cookie)
}

func (m *Manager) GC() {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.provider.SessionGC(m.maxlifetime)
	time.AfterFunc(time.Duration(m.maxlifetime), func() { m.GC() })
}

// Register Providers
func Register(name string, provider Provider) {
	if provider == nil {
		panic("Provider is nil")
	}
	if _, ok := providers[name]; ok {
		panic("Double provider")
	}
	providers[name] = provider
}

// init func
func init() {
	GlobalManager, _ = NewSessionManager("p", "m", 3600) // 1小时过期
}

type SessionStore struct {
	sid        string
	timeAccept time.Time
	value      map[interface{}]interface{} //  session的值的存储
}

type ProviderImp struct {
	lock     sync.RWMutex
	sessions map[string]*list.Element
	list     *list.List
}

func (pi *ProviderImp) SessionInit(sid string) (Session, error) {
	pi.lock.Lock()
	defer pi.lock.Unlock()
	newSess := &SessionStore{sid: sid, timeAccept: time.Now(), value: make(map[interface{}]interface{}, 0)}
	ele := pi.list.PushBack(newSess)
	pi.sessions[sid] = ele
	return newSess, nil
}

func (pi *ProviderImp) SessionRead(sid string) (Session, error) {
	pi.lock.Lock()
	defer pi.lock.Unlock()
	if ele, ok := pi.sessions[sid]; ok {
		return ele.Value.(*SessionStore), nil
	} else {
		sess, err := pi.SessionInit(sid)
		return sess, err
	}
}

func (pi *ProviderImp) SessionDestory(sid string) error {
	pi.lock.Lock()
	defer pi.lock.Unlock()
	if ele, ok := pi.sessions[sid]; ok {
		delete(pi.sessions, sid)
		pi.list.Remove(ele)
		return nil
	}
	return nil
}

func (pi *ProviderImp) SessionGC(maxlifetime int64) {
	pi.lock.Lock()
	defer pi.lock.Unlock()
	for {
		ele := pi.list.Back()
		if ele == nil {
			break
		}
		// 修剪
		if ele.Value.(*SessionStore).timeAccept.Unix()+maxlifetime >= time.Now().Unix() {
			pi.list.Remove(ele)
			delete(pi.sessions, ele.Value.(*SessionStore).sid)
		} else {
			break
		}
	}
}

func (pi *ProviderImp) SessionUpdate(sid string) error {
	pi.lock.Lock()
	defer pi.lock.Unlock()
	if ele, ok := pi.sessions[sid]; ok {
		ele.Value.(*SessionStore).timeAccept = time.Now()
		pi.list.MoveToFront(ele)
		return nil
	}
	return nil
}

func (ss *SessionStore) Set(key, value interface{}) error {
	ss.value[key] = value
	return pi.SessionUpdate(ss.sid)
}

func (ss *SessionStore) Get(key interface{}) interface{} {
	pi.SessionUpdate(ss.sid)
	if ele, ok := ss.value[key]; ok {
		return ele
	} else {
		return nil
	}
}

func (ss *SessionStore) Delete(key interface{}) error {
	pi.SessionUpdate(ss.sid)
	delete(ss.value, key)
}
func (ss *SessionStore) SessionID() string {
	return ss.sid
}
