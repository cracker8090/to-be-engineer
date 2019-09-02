



## Ionic、Angular、Cordova



1.即使我们将移动端web页面做得和原生应用及其相似，在我们的页面中也无法像原生应用那样调用原生的能力，当然通过输入框触发键盘、图库、拍照等操作不在这里“调用原生能力”的范畴。

2.单纯的web页面不能提交到应用商店被用户使用

## Ionic和Angular

Ionic是Angular的衍生品，Angular是单独的JS库，和jQuery一样能够独立用于开发应用，而Ionic只是对Angular进行了扩展，利用Angular实现了很多符合移动端应用的组件，并搭建了很完善的样式库，是对Angular最成功的应用样例。即使不使用Ionic，Angular也可与任意样式库，如Bootstrap、Foundation等搭配使用，得到想要的页面效果。

## Ionic/Angular和Cordova

“Cordova比Ionic/Angular好吗？”它们在混合开发中扮演的是不同的角色–Ionic/Angular负责页面的实现，而Cordova负责将实现的页面包装成原生应用（Android:apk；iOS:ipa）。就像花生，最内层的花生仁是Angular，花生仁的表皮是Ionic，而最外层的花生壳则是Cordova。包装完成之后我们的页面才有可能调用设备的原生能力，最后才能上传到应用商店被用户使用。

## Ionic/Angular和Cordova插件

关于Cordova插件要明确以下几点：

    Cordova插件的作用是提供一个桥梁供页面和原生通信，首先我们的页面不能直接调用设备能力，所以需要与能够调用设备能力的原生代码（Android:Java；iOS:OC）通信，此时就需要Cordova插件了。
    
    Cordova插件能够再任何Cordova工程中使用，和使用什么前端框架（如Ionic）无关。
    
    Ionic 2中封装了Ionic Native，方便了Cordova插件的使用，但在Ionic 2中仍然可以像Ionic 1中一样使用Cordova插件，Ionic Native不是必须的。
    
    即使在Ionic 2中使用了Ionic Native，也首先需要手动添加插件，如：cordova plugin add cordova-plugin-pluginName。


































