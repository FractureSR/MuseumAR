// pages/ticketDetail/ticketDetail.js
Page({

  /**
   * 页面的初始数据
   */
  data: {
  info_1: "1. 景区实行实名制，请下单时谨慎检查录入真实姓名和身份证号码，避免因个人原因录入身份证号码错误而导致无法入园，仅支持中国大陆手机号码订购；",
  info_2:"2. 身份证限购：一个身份证出行当天仅限订购一张门票（针对全部售票渠道限购），重复下单无效,请勿重复下单购买",
  info_3:"3. 根据疫情防控需要，景区单日接待量不超过8000人，超过上限将采取临时闭馆措施，游客参观需全程佩戴口罩，与他人距离保持1.5米以上；"
  },
  popConfirm: function(){
    wx.showModal({
      title: '确认购买',
      content: this.data.tname+'\n'+'参观日期 '+this.data.time+'\n'+'付款         ￥'+this.data.price,
      success: function (res) {
        if (res.confirm) {  
          console.log('点击确认回调')
        } else {   
          console.log('点击取消回调')
        }
      }
    })
  },
  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
  let That = this;
  That.setData({
    tname: options.tname,
    time: options.time,
    price:options.price
  })
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