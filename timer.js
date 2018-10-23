/* 在小程序中做倒计时 */
Page({
    data: {
        seconds: 60
    },
    countDown: function(){
        let that = this;
        that.setData({
            seconds: 60
        });
        let seconds = that.data.seconds;
        that.setData({
            timer: setInterval(function(){
                seconds --;
                that.setData({
                    seconds: seconds
                })
                if(seconds == 0){
                    clearInterval(that.data.timer);
                }
            }, 1000)
        });
    }
})
