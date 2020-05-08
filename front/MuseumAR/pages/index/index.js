//index.js
const app = getApp()

Page({
  data: {
    notice_pic:'',
    index_request_url:'',
    token:'init_token',
    current: '博物馆简介',
    current_scroll: '博物馆简介',

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

    //判断是用户是否绑定了
    if (app.globalData.Token && app.globalData.Token!= '') {
      this.setData({
      });
    } else {
      // 由于 getUserInfo 是网络请求，可能会在 Page.onLoad 之后才返回
      // 所以此处加入 callback 以防止这种情况
      app.tokenReadyCallback =() => {
        console.log('token!:: '+app.globalData.Token)
          this.setData({
            token:app.globalData.Token,
            index_request_url:'https://www.fracturesr.xyz/museumar/authenticated/museumhomepage?museum=000001&token=' + app.globalData.Token
          });  
          console.log("???: "+this.data.index_request_url)
        //向数据库请求获取博物馆数据
        wx.request({
          url: this.data.index_request_url,//json数据地址
          headers: {
            'Content-Type': 'application/json'
          },
          success: function (res) {
            console.log(res.data)
            console.log('Connected_index')
            _this.setData({
              //museum.id: res.data.Museum.Name,
              museum:res.data.Museum,
              sections:res.data.Sections,
              notices:res.data.Notices

            })
            //图片src处理
            let temp=[]
            _this.data.museum.Pictures.forEach(function (item, index) {
              temp.push('https://'+item.replace('/museumar',''));
            

            })
            //暂时写死
            let temp0 = 'https://'+ _this.data.notices[0].Picture.replace('/museumar', '')
            console.log("temp0:"+temp0)
            _this.setData({
              imgList:temp,
              notice_pic:temp0
            })
          },
          fail: function () {
            console.log('fail_index:' + _this.data.index_request_url)
          }
        })
      }
      
    }
    

    /*
    let temp=[];
    for (let i=1;i<=5;i++){
      temp.push("/source/index_"+i+".jpg");
    }
    this.setData({
      imgList:temp
    })
    console.log("GGGGG"+app.globalData.Token)
    */
  },
  //页面跳转函数
  ToSectionPage :function(e){
    wx.navigateTo({
      url: '/pages/SectionPage/SectionPage',
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
  ToScanPage: function(e){
    wx.navigateTo({
      url: '/pages/ScanPage/ScanPage',
    })
  },


})
