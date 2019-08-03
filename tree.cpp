#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

struct Node {
    int x, y;
};

bool compares(Node &n1, Node &n2) {
    if(n1.x == n2.x) {
        return n1.y < n2.y;
    }
    return n1.x < n2.x;
}

void read_node(Node& node) {
    int x, y;
    cin >> x >> y;
    if( x<0 ) {
        x = 0;
    }
    if(y < x) {
        swap(x, y);
    }
    node.x = x, node.y = y;
}

int main()
{
    int road, n;
    vector<Node> vec;
    while(cin >> road >> n) {
        vec.clear();
        for( int i=0; i<n; i++ ){
            Node node;
            read_node(node);
            vec.push_back(node);
        }
        sort(vec.begin(), vec.end(), compares);
        vector<Node>::iterator iter = vec.begin()+1;
        int res = vec[0].y-vec[0].x+1;
        Node now = vec[0];
        while(iter != vec.end()) {
            if((*iter).x > now.y) {
                res += ((*iter).y - (*iter).x)+1;
                 now = (*iter);
            } else if ((*iter).y > now.y) {
                res += ((*iter).y - now.y);
                now.y = (*iter).y;
            }
            iter++;
        }
        cout << road - res + 1 << endl;
    }
    return 0;
}
