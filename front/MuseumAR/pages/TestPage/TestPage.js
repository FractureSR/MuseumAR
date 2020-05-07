// pages/TestPage/TestPage.js
const { $Message } = require('../../components/dist/base/index');

Page({

  /**
   * 页面的初始数据
   */
  data: {
    visible1: false,
    visible2: false,
    visible3: false,
    visible4: false,
    visible5: false,
    actions3: [
      {
        name: '现金支付',
        color: '#2d8cf0',
      },
      {
        name: '微信支付',
        color: '#19be6b'
      },
      {
        name: '取消'
      }
    ],
    actions4: [
      {
        name: '按钮1'
      },
      {
        name: '按钮2',
        color: '#ff9900'
      },
      {
        name: '按钮3',
        icon: 'search'
      }
    ],
    actions5: [
      {
        name: '取消'
      },
      {
        name: '删除',
        color: '#ed3f14',
        loading: false
      }
    ]
  },
  handleOpen1() {
    this.setData({
      visible1: true
    });
  },

  handleClose1() {
    this.setData({
      visible1: false
    });
  },

  handleOpen2() {
    this.setData({
      visible2: true
    });
  },

  handleClose2() {
    this.setData({
      visible2: false
    });
  },

  handleOpen3() {
    this.setData({
      visible3: true
    });
  },

  handleClick3({ detail }) {
    const index = detail.index;

    if (index === 0) {
      $Message({
        content: '点击了现金支付'
      });
    } else if (index === 1) {
      $Message({
        content: '点击了微信支付'
      });
    }

    this.setData({
      visible3: false
    });
  },

  handleOpen4() {
    this.setData({
      visible4: true
    });
  },

  handleClick4() {
    this.setData({
      visible4: false
    });
  },

  handleOpen5() {
    this.setData({
      visible5: true
    });
  },

  handleClick5({ detail }) {
    if (detail.index === 0) {
      this.setData({
        visible5: false
      });
    } else {
      const action = [...this.data.actions5];
      action[1].loading = true;

      this.setData({
        actions5: action
      });

      setTimeout(() => {
        action[1].loading = false;
        this.setData({
          visible5: false,
          actions5: action
        });
        $Message({
          content: '删除成功！',
          type: 'success'
        });
      }, 2000);
    }
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {

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