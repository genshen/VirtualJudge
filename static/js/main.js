/**
 * Created by gensh on 2016/11/3.
 */
const Config = {
    base: "/",
    appName: "Virtual Judge",
    logo: "https://v4-alpha.getbootstrap.com/assets/brand/bootstrap-solid.svg",
    auth: {
        github: {
            name: "Github",
            logo: "/static/img/github.png",
            enable: true,
            client_id: "94b2759733e24ee2994d",
            redirect_uri: "http://localhost:8080/auth/callback/github",
            auth_url: "https://github.com/login/oauth/authorize/?redirect_uri=R&client_id=C",
        }
    }
};
var userInfo = {name: "", avatar: "", id: 0, is_login: false};

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
        return {Problems: []};
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
        goSignIn: function (key) { //eg:go to github.com
            var instance = this.config.auth[key];
            var url = instance.auth_url.replace("R", instance.redirect_uri).replace("C", instance.client_id);
            // console.log(url);
            window.open(url, "", "location=no,status=no");
            $("#sign-in-modal").modal("hide");
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
                        userInfo.avatar = data.avatar;
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

window.addEventListener('message', function (e) {
    if (e.origin == location.origin) {
        var data = e.data;
        if (data.status == 1) {
            userInfo.name = data.name;
            userInfo.id = data.id;
            userInfo.avatar = data.avatar;
            userInfo.is_login = true;
            new Snackbar("登录认证成功", {timeout: 3500});
        }
    }
});