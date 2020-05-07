//app.js
App({
  globalData: {
    userInfo: null,
    Token: []
  },
  Data: {
    Code: [],
    request_url: []
  },
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
        this.Data.Code = res.code;
        this.Data.request_url = 'https://www.fracturesr.xyz/museumar/signin?code='+this.Data.Code;
        this.globalData.Token=this.Data.request_url;
        //console.log(this.Data.request_url);
        
        //请求与数据库通信
        wx.request({
          url: getApp().Data.request_url,//json数据地址
          success: function (res) {
            console.log(1)
            console.log(res.data)
             this.setData({
               Token:res.data.token,
               //res代表success函数的事件对，data是固定的，imgListData是上面json数据中imgListData
             })
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

    
    
  },
  
})