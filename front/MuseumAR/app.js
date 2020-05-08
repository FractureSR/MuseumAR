//app.js
App({
  
  onLaunch: function () {
    // 展示本地存储能力
    var logs = wx.getStorageSync('logs') || []
    logs.unshift(Date.now())
    wx.setStorageSync('logs', logs)

    // 登录
    wx.login({
      success: res => {
        // 发送 res.code 到后台换取 openId, sessionKey, unionId
        console.log(res.code);
        this.globalData.Code = res.code;
        this.globalData.request_url = 'https://www.fracturesr.xyz/museumar/signin?code='+this.globalData.Code;
        console.log(this.globalData.request_url)
        console.log("url: "+ getApp().globalData.request_url)

        //请求与数据库通信
        wx.request({
          url: getApp().globalData.request_url,//json数据地址
          headers: {
            'Content-Type': 'application/json'
          },
          success: function (res) {
            console.log("Success")
            getApp().globalData.Token=res.data.token  
            console.log("token::"+getApp().globalData.Token)
            //由于这里是网络请求，可能会在 Page.onLoad 之后才返回
            // 所以此处加入 callback 以防止这种情况
           
            if (getApp().tokenReadyCallback) {
              getApp().tokenReadyCallback();
              console.log("app.js:tokenReadyCallback")
            }
            
            
          },
          fail: function () {
            console.log('!!!'+getApp().Data.request_url);
          }
        })
        
      }
    })
    // 获取用户信息
    wx.getSetting({
      success: res => {
        if (res.authSetting['scope.userInfo']) {
          // 已经授权，可以直接调用 getUserInfo 获取头像昵称，不会弹框
          wx.getUserInfo({
            success: res => {
              // 可以将 res 发送给后台解码出 unionId
              this.globalData.userInfo = res.userInfo

              // 由于 getUserInfo 是网络请求，可能会在 Page.onLoad 之后才返回
              // 所以此处加入 callback 以防止这种情况
              if (this.userInfoReadyCallback) {
                this.userInfoReadyCallback(res)
              }
            }
          })
        }
      }
    })

    
    
  }, globalData: {
    userInfo: null,
    Token: [],
    Code: [],
    request_url: [],
    employId: ''

  }
})