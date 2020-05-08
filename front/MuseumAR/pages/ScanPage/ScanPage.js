// pages/ScanPage/ScanPage.js
Page({

  /**
   * 页面的初始数据
   */
  data: {

  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function () {
 
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

  },
/**
 * 不断拍照向服务器上传图片，成功则返回展品图片和介绍，失败则继续拍照上传
 */
  Camera_initdone: function(){
    console.log("Camera_init_done!")
    const cameraContext = wx.createCameraContext();
    for (let check=false;check== true;){
      cameraContext.takePhoto({
        quality: "high", //高质量的图片
        success: res => {
          //res.tempImagePath照片文件在手机内的的临时路径
          let tempImagePath = res.tempImagePath
          wx.uploadFile({
            url: 'https://example.weixin.qq.com/upload', //仅为示例，非真实的接口地址  
            filePath: tempFilePath,
            name: 'file',
            formData: {
              //formData:HTTP 请求中其他额外的参数
              'user': 'test'
            },
            success: (res) => {
              //上传成功  
              this.setData({
                check:true,//不再上传图片
                //返回图片和文字
                img:'' ,//展品图片
                name:'' ,//展品名字
                introduction: ''//展品介绍
              })
              wx.navigateTo({
                url: '/pages/ARDetailPage/ARDetailPage',
                data:''
              })

            },
            fail: function (t) {
              //上传失败  
            },
          })
        }
      })
    }
  }
})