// miniprogram/pages/StoreDetailPage/StoreDetailPage.js
const { $Message } = require('../../components/dist/base/index');
Page({

  /**
   * 页面的初始数据
   */
  data: {
   
    buy_confirm_visiable:false,
    buy_confirm_actions: [
      {
        name: '取消',
        color: '#2d8cf0',
      },
      {
        name: '确认',
        color: '#19be6b'
      },
    ],
    //商品详情页标签转换
    tab_current: '1',
    tab_current_scroll: '1',
    //swiper变量
    hidden: false,//
    imgList: [],// 图片src数组
    autoplay: false,//控制自动播放
    circular: true,//循环播放
    indicatordots: true,//控制底下显示的圆点
    duration: 500, //滑动动画时长
    interval: 5000 //如果开启自动播放，控制切换时间间隔。
  },
  ////显示确认购买弹窗
  Buy_Confirm_Modal_Open(){
    wx.showModal({
      title: '确认购买',
      content: '商品名称: '+this.data.title+'\n'  + '应付金额: ￥'+this.data.price,
      success: function (res) {
        if (res.confirm) {
          $Message({
            content: '购买成功！'
          });
        } else {
        }
      }
    })
  },
  ////确认购买弹窗
  Buy_Confirm({ detail }) {
    console.log(detail.index)
    if (detail.index === 1) {
      $Message({
        content: '购买成功！'
      });
    } 
    this.setData({
      buy_confirm_visiable: false
    });
  },
  ////swiper切换函数
  onSlideChange: function (event) {
    var postId = event.detail.current;
    console.log(postId);
  },
  //标签变换
  Tab_HandleChange({ detail }) {
    this.setData({
      tab_current: detail.key,
      
    });
  },
  Tab_HandleChangeScroll({ detail }) {
    this.setData({
      tab_current_scroll: detail.key
    });
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
    /*
    let temp = [];
    for (let i = 1; i <= 4; i++) {
      temp.push("/pages/StorePage/" + i + ".jpg");
    }
    this.setData({
      imgList: temp
    })
    */
    let temp=[];
    temp.push(options.image_src);
    console.log(options);
    this.setData({
      
      //以下为storedetial含有内容
      image_src:options.image_src,//image_src 图片src，目前只能传一张图片进来
      introduction:options.introduction,//introduction 简介
      price:options.price, //price 价格
      title:options.title,//title 标题
      imgList:temp
      //后期应做成只传商品ID，然后该页面向服务器请求信息
      
    })
    //console.log(storedetail)
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