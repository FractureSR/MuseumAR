//index.js
const app = getApp()

Page({
  data: {
    current: '博物馆简介',
    current_scroll: '博物馆简介',
    museum:{
      id:'000001',
      name:'',
      description:'',
      pictures:[]
    },
    //swiper变量
    hidden: false,//
    imgList: [],// 图片src数组
    autoplay: true,//控制自动播放
    circular: true,//循环播放
    indicatordots: true,//控制底下显示的圆点
    duration: 500, //滑动动画时长
    interval: 5000 //如果开启自动播放，控制切换时间间隔。
  },

  ////swiper切换函数
  onSlideChange: function (event) {
    var postId = event.detail.current;
    console.log(postId);
  },

  ////标签页相关转换函数
  handleChange({ detail }) {
    this.setData({
      current: detail.key
    });
  },
  handleChangeScroll({ detail }) {
    this.setData({
      current_scroll: detail.key
    });
  },
  
  onLoad: function() {
    var _this = this
    /*wx.request({
      url: '',//json数据地址
      headers: {
        'Content-Type': 'application/json'
      },
      success: function (res) {
        console.log(res.data)
        _this.setData({
          museum:res.data.Museum,
          //res代表success函数的事件对，data是固定的，imgListData是上面json数据中imgListData
        })
      }
    })*/
    let temp=[];
    for (let i=1;i<=4;i++){
      temp.push("/pages/StorePage/"+i+".jpg");
    }
    this.setData({
      imgList:temp
    })
  },
  //页面跳转函数
  ToImformationPage :function(e){
    wx.navigateTo({
      url: '',
    })
  },
  ToTicketPage: function (e) {
    wx.navigateTo({
      url: '/pages/TicketPage/TicketPage',
    })
  },
  ToStorePage: function (e) {
    wx.navigateTo({
      url: '/pages/StorePage/StorePage',
    })
  },
  ToCommentPage: function (e) {
    wx.navigateTo({
      url: '/pages/CommentPage/CommentPage',
    })
  },





})
