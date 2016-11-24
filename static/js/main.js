/**
 * Created by gensh on 2016/11/3.
 */
const Config = {
    base: "/",
    appName: "Virtual Judge",
    logo: "https://v4-alpha.getbootstrap.com/assets/brand/bootstrap-solid.svg",
    auth: {
        github: {
            enable: true,
            authCallback: "/auth/github_auth/callback"
        }
    }
    // authUrl: "https://graph.qq.com/oauth/show?which=ConfirmPage&client_id=101363082&response_type=code&scope=get_user_info&state=10.16.55.200"
};
var userInfo = {name: "", id: 0, is_login: false};

//init config
// Config.authUrl += "&redirect_uri=" + Config.authCallback;

var Home = Vue.extend({
    template: '#template-home',
    created: function () {
        var toolbar = $('#vj_app_toolbar');
        var docToolbarHeight = toolbar.outerHeight();
        $(window).on('scroll', function () {
            if ($(this).scrollTop() > docToolbarHeight) {
                toolbar.addClass('waterfall');
            } else {
                toolbar.removeClass('waterfall');
            }
        });
    },
    mounted: function () {
        $('#vj_app_toolbar').removeClass('waterfall');
    },
    destroyed: function () {
        $(window).off('scroll');
        $('#vj_app_toolbar').addClass('waterfall');
    }
});

var Problems = Vue.extend({
    template: '#template-problems',
    data: function () {
        return {userInfo: userInfo, Problems: []};
    },
    methods: {},
    created: function () {

    }
});

const router = new VueRouter({
    base: Config.base,
    routes: [
        {path: '/', name: 'home', component: Home},
        {path: '/problems', name: 'problems', component: Problems}
    ]
});

var app = new Vue({
    router: router,
    data: {config: Config},
    methods: {
        qqLogin: function () {
            window.open(Config.authUrl, "Github", "status=no,titlebar=no,toolbar=no,menubar=no");
            console.log("s");
        }
    },
    mounted: function () {
        $.ajax({
            url: Config.base + "auth/user_status",
            success: function (data) {
                try {
                    if (data.is_login) {
                        userInfo.name = data.name;
                        userInfo.id = data.id;
                        userInfo.is_login = true;
                    }
                } catch (e) {
                    new Snackbar("加载数据出错啦,请刷新重试.", {timeout: 3500});
                }
            }, error: function (err) {
                new Snackbar("加载数据出错啦,请刷新重试.", {timeout: 3500});
            }
        });
    }
}).$mount('#app');
//
// setTimeout(function () {
//     console.log(app.$refs.app.onLoginSuccess());
// }, 2000);