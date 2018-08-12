import sys
from PyQt5.QtWidgets import (QWidget, QProgressBar,
                             QPushButton, QApplication)
from PyQt5.QtCore import QBasicTimer


class Example(QWidget):

    def __init__(self):
        super(Example, self).__init__()
        self.init_ui()

    def init_ui(self):
        self.bar = QProgressBar(self)
        self.bar.setGeometry(30, 40, 200, 25)
        
        self.btn = QPushButton('Start', self)
        self.btn.move(40, 80)
        self.btn.clicked.connect(self.do_action)

        self.timer = QBasicTimer()
        self.step = 0

        self.setGeometry(300, 200, 250, 170)
        self.setWindowTitle('Process Bar')
        self.show()

    def timerEvent(self, event):
        if self.step > 100:
            self.timer.stop()
            self.btn.setText('Finished')
        else:
            self.step += 1
            self.bar.setValue(self.step)

    def do_action(self):
        sender = self.sender()
        if sender.text() == 'Finished':
            self.close()
        if self.timer.isActive():
            self.timer.stop()
            if sender.text() == 'Pause':
                self.btn.setText('Continue')
        else:
            self.timer.start(100, self)
            if sender.text() == 'Continue':
                self.btn.setText('Pause')
            else:
                self.btn.setText('Pause')

if __name__ == '__main__':
    app = QApplication(sys.argv)
    ex = Example()
    sys.exit(app.exec_())
