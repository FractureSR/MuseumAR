// pages/TestPage/TestPage.js
const { $Message } = require('../../components/dist/base/index');
const  app=getApp();
Page({

  /**
   * 页面的初始数据
   */
  data: {
    
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
    var _this = this

    //判断是用户是否绑定了
    if (app.globalData.Token && app.globalData.Token != '') {
      this.setData({
        token: app.globalData.Token,
        index_request_url: 'https://www.fracturesr.xyz/museumar/authenticated/museumhomepage?museum=000001&token=' + app.globalData.Token
      });
       wx.request({
          url: this.data.index_request_url,//json数据地址
          headers: {
            'Content-Type': 'application/json'
          },
          success: function (res) {
            console.log(res.data)
            console.log('Connected_SectionPage')
            _this.setData({
              //museum.id: res.data.Museum.Name,
              museum: res.data.Museum,
              sections: res.data.Sections,
              notices: res.data.Notices

            })
            //图片src处理
          },
          fail: function () {
            console.log('fail_index:' + _this.data.index_request_url)
          }
        })
    } else {
      // 由于 getUserInfo 是网络请求，可能会在 Page.onLoad 之后才返回
      // 所以此处加入 callback 以防止这种情况
      app.tokenReadyCallback = () => {
        console.log('token!:: ' + app.globalData.Token)
        this.setData({
          token: app.globalData.Token,
          index_request_url: 'https://www.fracturesr.xyz/museumar/authenticated/museumhomepage?museum=000001&token=' + app.globalData.Token
        });
        console.log("???: " + this.data.index_request_url)
        //向数据库请求获取博物馆数据
        wx.request({
          url: this.data.index_request_url,//json数据地址
          headers: {
            'Content-Type': 'application/json'
          },
          success: function (res) {
            console.log(res.data)
            console.log('Connected_SectionPage')
            _this.setData({
              //museum.id: res.data.Museum.Name,
              museum: res.data.Museum,
              sections: res.data.Sections,
              notices: res.data.Notices

            })
            //图片src处理
          },
          fail: function () {
            console.log('fail_index:' + _this.data.index_request_url)
          }
        })
      }

    }
  },

  /**
   * 生命周期函数--监听页面初次渲染完成
   */
  onReady: function () {

  },

  /**
   * 生命周期函数--监听页面显示
   */
  onShow: function () {

  },

  /**
   * 生命周期函数--监听页面隐藏
   */
  onHide: function () {

  },

  /**
   * 生命周期函数--监听页面卸载
   */
  onUnload: function () {

  },

  /**
   * 页面相关事件处理函数--监听用户下拉动作
   */
  onPullDownRefresh: function () {

  },

  /**
   * 页面上拉触底事件的处理函数
   */
  onReachBottom: function () {

  },

  /**
   * 用户点击右上角分享
   */
  onShareAppMessage: function () {

  }
})