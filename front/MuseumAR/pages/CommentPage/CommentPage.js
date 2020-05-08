
const app = getApp()

Page({
    data: {
      talks: [],
      inputValue: '',
      inputBiaoqing: '',
      faces: [ 'http://pic2.zhimg.com/50/v2-958d33fd4a4de747058adcafdf753074_hd.jpg'],
      names: ['贝贝'],
    },
    //解决滑动穿透问题
    emojiScroll: function(e) {
      console.log(e)
    },
    onReady: function() {
      // 评论弹出层动画创建
      this.animation = wx.createAnimation({
        duration: 400, // 整个动画过程花费的时间，单位为毫秒
        timingFunction: "ease", // 动画的类型
        delay: 0 // 动画延迟参数
      })
    },
    showTalks: function() {
      // 加载数据
      this.loadTalks();
      // 设置动画内容为：使用绝对定位显示区域，高度变为100%
      this.animation.bottom("0rpx").height("100%").step()
      this.setData({
        talksAnimationData: this.animation.export()
      })
    },
   
    hideTalks: function() {
      // 设置动画内容为：使用绝对定位隐藏整个区域，高度变为0
      this.animation.bottom("-100%").height("0rpx").step()
      this.setData({
        talks: [],
        talksAnimationData: this.animation.export()
      })
    },
    onScrollLoad: function() {
      // 加载新的数据
      this.loadTalks();
    },
    //下拉评论框隐藏
    touchStart: function(e) {
      let touchStart = e.touches[0].clientY;
      this.setData({
        touchStart,
      })
    },
    touchMove: function(e) {
      // console.log(this.data.touchStart)
      let touchLength = e.touches[0].clientY - this.data.touchStart;
      console.log(touchLength - 100)
      if (touchLength > 100) {
        this.animation.bottom("-100%").height("0rpx").step()
        this.setData({
          talks: [],
          talksAnimationData: this.animation.export(),
        })
      }
    },
    //输入框失去焦点时触发
    bindInputBlur: function(e) {
      console.log(e)
      console.log(this.data.inputBiaoqing)
      this.data.inputValue = e.detail.value + this.data.inputBiaoqing;
    },
  //授权登入
 
  bindGetUserInfo(e) {
    console.log(e.detail.userInfo)
  },
 //请求与数据库通信

    //点击发布，发布评论
    faBu: function() {
    //授权得到头像和nickname
      let that = this;

     wx.getSetting({
    success(res) {
      if (res.authSetting['scope.userInfo']) {
        // 已经授权，可以直接调用 getUserInfo 获取头像昵称
        wx.getUserInfo({
          success: function (res) {
            console.log(res.userInfo),
              that.setData({
                nickName: res.userInfo.nickName,
                avatarUrl: res.userInfo.avatarUrl,
              })
          }
        })
      }
    }
  })
  //请求与数据库通信
  
  //连接数据库
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
          token:app.globalData.Token
        });     
    }
  }
  
  //得到内容
      this.data.talks.unshift({
       // avatarUrl: this.data.faces[Math.floor(Math.random() * this.data.faces.length)],
        //nickName: this.data.names[Math.floor(Math.random() * this.data.names.length)],
        content: this.data.inputValue,
        talkTime: Date()
      })
      that.data.inputValue = '';
      that.setData({
        talks: that.data.talks,
        inputValue: that.data.inputValue,
        talksAnimationData: that.animation.export()
      })
   
    },
  //删除
  deleteComment:function(e){
  
    let that = this
    let index = e.target.dataset.index
    let arrayLength = that.data.talks.length 
    let newArray = []
    if (arrayLength > 0) {
      for (let i = 0; i <arrayLength; i++) {
        if (i !== index) {
          newArray.push(that.data.talks[i])
        }
      }
      that.setData({
       talks: newArray
      })
    } else {
      wx.showToast({
        icon: 'none',
        title: '最后一条不可删除',
      })
    }
  },
 
  /**
   * 上传图片方法
   */
  upload: function () {
    let that = this;
    wx.chooseImage({
      count: 9, // 默认9
      sizeType: ['original', 'compressed'], // 可以指定是原图还是压缩图，默认二者都有
      sourceType: ['album', 'camera'], // 可以指定来源是相册还是相机，默认二者都有
      success: res => {
        wx.showToast({
          title: '正在上传...',
          icon: 'loading',
          mask: true,
          duration: 1000
        })
        // 返回选定照片的本地文件路径列表，tempFilePath可以作为img标签的src属性显示图片
        let tempFilePaths = res.tempFilePaths;

        that.setData({
          tempFilePaths: tempFilePaths
        })
       
        }

      }
    )
  },
  /**
   * 预览图片方法
   */
  listenerButtonPreviewImage: function (e) {
    let index = e.target.dataset.index;
    let that = this;
    console.log(that.data.tempFilePaths[index]);
    console.log(that.data.tempFilePaths);
    wx.previewImage({
      current: that.data.tempFilePaths[index],
      urls: that.data.tempFilePaths,
      //这根本就不走
      success: function (res) {
        //console.log(res);
      },
      //也根本不走
      fail: function () {
        //console.log('fail')
      }
    })
  },
  /**
   * 长按删除图片
   */
  deleteImage: function (e) {
    var that = this;
    var tempFilePaths = that.data.tempFilePaths;
    var index = e.currentTarget.dataset.index;//获取当前长按图片下标
    wx.showModal({
      title: '提示',
      content: '确定要删除此图片吗？',
      success: function (res) {
        if (res.confirm) {
          console.log('点击确定了');
          tempFilePaths.splice(index, 1);
        } else if (res.cancel) {
          console.log('点击取消了');
          return false;
        }
        that.setData({
          tempFilePaths
        });
      }
    })
  }

  })
  