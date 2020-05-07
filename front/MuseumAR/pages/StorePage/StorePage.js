// miniprogram/pages/StorePage/StorePage.js
Page({

  /**
   * 页面的初始数据
   */
  data: {
    storelist: [],
    imgList: [],// 图片src数组
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
    
    
    let temp = [];
    let image_src=[];
    let introduction=[];
    let price=[];
    let title =[];
    let index=[];
    for (let i = 1; i <= 4; i++) {
      image_src = "/pages/StorePage/" + i + ".jpg";
      title = "商品" + i;
      price =i+"88";
      introduction= "这是商品"+i;
      index = i;
      temp.push({index,image_src, title,introduction, price});
    };
    this.setData({
      storelist:temp
    })
  },
  //跳转至商品详情页
  ToStoreDetailPage: function(event){
    //console.log(event.currentTarget.dataset.item) //获取点击的商品项
    let _item = event.currentTarget.dataset.item
    wx.navigateTo({
      url: '/pages/StoreDetailPage/StoreDetailPage?price='+_item.price+'&title='+_item.title+'&image_src='
+_item.image_src+'&introduction='+_item.introduction
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