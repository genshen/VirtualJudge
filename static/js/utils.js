function Snackbar(t,e,i){""!==t?(this.options=this.activateOptions(e),this.data=t,this.callback=i,this.start(),this.snackbar()):console.warn("SnackbarLight: You can not create a empty snackbar please give it a string.")}Snackbar.prototype={options:{timeout:5e3,activeClass:"active",link:!1,url:"#"},start:function(){if(!document.getElementById("snackbar-container")){var t=document.createElement("div");t.setAttribute("id","snackbar-container"),document.body.appendChild(t)}},timer:function(t,e){var i=e;this.timer={timerId:Math.round(1e3*Math.random()),pause:function(){window.clearTimeout(this.timerId),i-=new Date-start},resume:function(){start=new Date,window.clearTimeout(this.timerId),this.timerId=window.setTimeout(t,i)}},this.timer.resume()},snackbar:function(){var t=this,e=document.createElement("div");document.getElementById("snackbar-container").appendChild(e),e.innerHTML=this.getData(),e.setAttribute("class","snackbar"),setTimeout(function(){e.setAttribute("class","snackbar "+t.options.activeClass)},50),this.options.timeout!==!1&&this.timer(function(){e.setAttribute("class","snackbar"),t.destroy(e)},this.options.timeout),this.listeners(e)},getData:function(){return this.options.link!==!1?"<span>"+this.data+"</span><a href='"+this.options.url+"'>"+this.options.link+"</a>":"<span>"+this.data+"</span>"},listeners:function(t){var e=this;t.addEventListener("click",function(){"function"==typeof e.callback&&e.callback(),t.setAttribute("class","snackbar"),e.destroy(t)}),t.addEventListener("mouseenter",function(){e.timer.pause()}),t.addEventListener("mouseout",function(){e.timer.resume()})},destroy:function(t){this.timer.pause(),setTimeout(function(){t.remove()},1e4)},activateOptions:function(t){var e=this,i=t||{};for(var n in this.options)e.options.hasOwnProperty(n)&&!i.hasOwnProperty(n)&&(i[n]=e.options[n]);return i}},SnackbarLight={install:function(t){var e=this;t.prototype.$snackbar={},t.prototype.$snackbar.create=function(t,i,n){e.create(t,i,n)}},create:function(t,e,i){new Snackbar(t,e,i)}},"object"==typeof exports?module.exports=SnackbarLight:"function"==typeof define&&define.amd?define([],function(){return SnackbarLight}):window.Vue&&Vue.use(SnackbarLight);for(var elements=document.querySelectorAll("[data-toggle=snackbar]"),i=elements.length-1;i>=0;i--)elements[i].addEventListener("click",function(){var t={};null!==this.getAttribute("data-link")&&(t.link=this.getAttribute("data-link")),null!==this.getAttribute("data-timeout")&&(t.timeout=this.getAttribute("data-timeout")),null!==this.getAttribute("data-activeClass")&&(t.activeClass=this.getAttribute("data-active-class")),this.getAttribute("data-url")&&(t.url=this.getAttribute("data-url")),new Snackbar(this.getAttribute("data-content"),t)});
// https://github.com/joostlawerman/SnackbarLightjs

/*util：
 * including:
 * 1.snackbar
 * 2.base64.js
 * 3.jquery.cookie.js
 * 4. form post utils
 * 5 format time
 * */

//base64.js
// https://github.com/dankogai/js-base64
(function(global){"use strict";var _Base64=global.Base64;var version="2.1.9";var buffer;if(typeof module!=="undefined"&&module.exports){try{buffer=require("buffer").Buffer}catch(err){}}var b64chars="ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/";var b64tab=function(bin){var t={};for(var i=0,l=bin.length;i<l;i++)t[bin.charAt(i)]=i;return t}(b64chars);var fromCharCode=String.fromCharCode;var cb_utob=function(c){if(c.length<2){var cc=c.charCodeAt(0);return cc<128?c:cc<2048?fromCharCode(192|cc>>>6)+fromCharCode(128|cc&63):fromCharCode(224|cc>>>12&15)+fromCharCode(128|cc>>>6&63)+fromCharCode(128|cc&63)}else{var cc=65536+(c.charCodeAt(0)-55296)*1024+(c.charCodeAt(1)-56320);return fromCharCode(240|cc>>>18&7)+fromCharCode(128|cc>>>12&63)+fromCharCode(128|cc>>>6&63)+fromCharCode(128|cc&63)}};var re_utob=/[\uD800-\uDBFF][\uDC00-\uDFFFF]|[^\x00-\x7F]/g;var utob=function(u){return u.replace(re_utob,cb_utob)};var cb_encode=function(ccc){var padlen=[0,2,1][ccc.length%3],ord=ccc.charCodeAt(0)<<16|(ccc.length>1?ccc.charCodeAt(1):0)<<8|(ccc.length>2?ccc.charCodeAt(2):0),chars=[b64chars.charAt(ord>>>18),b64chars.charAt(ord>>>12&63),padlen>=2?"=":b64chars.charAt(ord>>>6&63),padlen>=1?"=":b64chars.charAt(ord&63)];return chars.join("")};var btoa=global.btoa?function(b){return global.btoa(b)}:function(b){return b.replace(/[\s\S]{1,3}/g,cb_encode)};var _encode=buffer?function(u){return(u.constructor===buffer.constructor?u:new buffer(u)).toString("base64")}:function(u){return btoa(utob(u))};var encode=function(u,urisafe){return!urisafe?_encode(String(u)):_encode(String(u)).replace(/[+\/]/g,function(m0){return m0=="+"?"-":"_"}).replace(/=/g,"")};var encodeURI=function(u){return encode(u,true)};var re_btou=new RegExp(["[À-ß][-¿]","[à-ï][-¿]{2}","[ð-÷][-¿]{3}"].join("|"),"g");var cb_btou=function(cccc){switch(cccc.length){case 4:var cp=(7&cccc.charCodeAt(0))<<18|(63&cccc.charCodeAt(1))<<12|(63&cccc.charCodeAt(2))<<6|63&cccc.charCodeAt(3),offset=cp-65536;return fromCharCode((offset>>>10)+55296)+fromCharCode((offset&1023)+56320);case 3:return fromCharCode((15&cccc.charCodeAt(0))<<12|(63&cccc.charCodeAt(1))<<6|63&cccc.charCodeAt(2));default:return fromCharCode((31&cccc.charCodeAt(0))<<6|63&cccc.charCodeAt(1))}};var btou=function(b){return b.replace(re_btou,cb_btou)};var cb_decode=function(cccc){var len=cccc.length,padlen=len%4,n=(len>0?b64tab[cccc.charAt(0)]<<18:0)|(len>1?b64tab[cccc.charAt(1)]<<12:0)|(len>2?b64tab[cccc.charAt(2)]<<6:0)|(len>3?b64tab[cccc.charAt(3)]:0),chars=[fromCharCode(n>>>16),fromCharCode(n>>>8&255),fromCharCode(n&255)];chars.length-=[0,0,2,1][padlen];return chars.join("")};var atob=global.atob?function(a){return global.atob(a)}:function(a){return a.replace(/[\s\S]{1,4}/g,cb_decode)};var _decode=buffer?function(a){return(a.constructor===buffer.constructor?a:new buffer(a,"base64")).toString()}:function(a){return btou(atob(a))};var decode=function(a){return _decode(String(a).replace(/[-_]/g,function(m0){return m0=="-"?"+":"/"}).replace(/[^A-Za-z0-9\+\/]/g,""))};var noConflict=function(){var Base64=global.Base64;global.Base64=_Base64;return Base64};global.Base64={VERSION:version,atob:atob,btoa:btoa,fromBase64:decode,toBase64:encode,utob:utob,encode:encode,encodeURI:encodeURI,btou:btou,decode:decode,noConflict:noConflict};if(typeof Object.defineProperty==="function"){var noEnum=function(v){return{value:v,enumerable:false,writable:true,configurable:true}};global.Base64.extendString=function(){Object.defineProperty(String.prototype,"fromBase64",noEnum(function(){return decode(this)}));Object.defineProperty(String.prototype,"toBase64",noEnum(function(urisafe){return encode(this,urisafe)}));Object.defineProperty(String.prototype,"toBase64URI",noEnum(function(){return encode(this,true)}))}} })(this);//if(global["Meteor"]){Base64=global.Base64}

//js-cookie
/*!
 * JavaScript Cookie v2.1.2
 * https://github.com/js-cookie/js-cookie
 *
 * Copyright 2006, 2015 Klaus Hartl & Fagner Brack
 * Released under the MIT license
 */
(function (factory) {
    if (typeof define === 'function' && define.amd) {
        define(factory);
    } else if (typeof exports === 'object') {
        module.exports = factory();
    } else {
        var OldCookies = window.Cookies;
        var api = window.Cookies = factory();
        api.noConflict = function () {
            window.Cookies = OldCookies;
            return api;
        };
    }
}(function () {
    function extend() {
        var i = 0;
        var result = {};
        for (; i < arguments.length; i++) {
            var attributes = arguments[i];
            for (var key in attributes) {
                result[key] = attributes[key];
            }
        }
        return result;
    }

    function init(converter) {
        function api(key, value, attributes) {
            var result;
            if (typeof document === 'undefined') {
                return;
            }
            // Write
            if (arguments.length > 1) {
                attributes = extend({
                    path: '/'
                }, api.defaults, attributes);

                if (typeof attributes.expires === 'number') {
                    var expires = new Date();
                    expires.setMilliseconds(expires.getMilliseconds() + attributes.expires * 864e+5);
                    attributes.expires = expires;
                }

                try {
                    result = JSON.stringify(value);
                    if (/^[\{\[]/.test(result)) {
                        value = result;
                    }
                } catch (e) {
                }

                if (!converter.write) {
                    value = encodeURIComponent(String(value))
                        .replace(/%(23|24|26|2B|3A|3C|3E|3D|2F|3F|40|5B|5D|5E|60|7B|7D|7C)/g, decodeURIComponent);
                } else {
                    value = converter.write(value, key);
                }

                key = encodeURIComponent(String(key));
                key = key.replace(/%(23|24|26|2B|5E|60|7C)/g, decodeURIComponent);
                key = key.replace(/[\(\)]/g, escape);

                return (document.cookie = [
                    key, '=', value,
                    attributes.expires && '; expires=' + attributes.expires.toUTCString(), // use expires attribute, max-age is not supported by IE
                    attributes.path && '; path=' + attributes.path,
                    attributes.domain && '; domain=' + attributes.domain,
                    attributes.secure ? '; secure' : ''
                ].join(''));
            }

            // Read

            if (!key) {
                result = {};
            }

            // To prevent the for loop in the first place assign an empty array
            // in case there are no cookies at all. Also prevents odd result when
            // calling "get()"
            var cookies = document.cookie ? document.cookie.split('; ') : [];
            var rdecode = /(%[0-9A-Z]{2})+/g;
            var i = 0;

            for (; i < cookies.length; i++) {
                var parts = cookies[i].split('=');
                var cookie = parts.slice(1).join('=');

                if (cookie.charAt(0) === '"') {
                    cookie = cookie.slice(1, -1);
                }

                try {
                    var name = parts[0].replace(rdecode, decodeURIComponent);
                    cookie = converter.read ?
                        converter.read(cookie, name) : converter(cookie, name) ||
                    cookie.replace(rdecode, decodeURIComponent);

                    if (this.json) {
                        try {
                            cookie = JSON.parse(cookie);
                        } catch (e) {
                        }
                    }

                    if (key === name) {
                        result = cookie;
                        break;
                    }

                    if (!key) {
                        result[name] = cookie;
                    }
                } catch (e) {
                }
            }

            return result;
        }

        api.set = api;
        api.get = function (key) {
            return api(key);
        };
        api.getJSON = function () {
            return api.apply({
                json: true
            }, [].slice.call(arguments));
        };
        api.defaults = {};

        api.remove = function (key, attributes) {
            api(key, '', extend(attributes, {
                expires: -1
            }));
        };

        api.withConverter = init;

        return api;
    }

    return init(function () {
    });
}));

var Util = {
    postData: {
        config: {
            authUrl: "" //todo
        },
        init: function (url, data, o, onPostSuccess, onUnAuth, onPostError, onError,onFinish) {
            var options = $.extend({}, {snackBarAlive: 4000,multiError:true, showNext: false, authUrl: this.config.authUrl}, o);
            if (!onError) {
                onError = function () {
                    new Snackbar("error happened,please<a href='" + window.location.href + "'>Refresh</a> and try again!",
                        {timeout: options.snackBarAlive});
                }
            }
            if (!onPostError) {
                if(options.multiError){
                    onPostError = function (Errors) {
                        for (var key in Errors) {
                            var err = Errors[key].Errors;
                            if (err.length > 0) {
                                new Snackbar(err[0].Message, {timeout: options.snackBarAlive});
                                return;
                            }
                        }
                    }
                }else{
                    onPostError = function (error) {
                        new Snackbar(error, {timeout: options.snackBarAlive});
                    }
                }
            }
            if (!onUnAuth) {
                onUnAuth = function () {
                    var url;
                    if (options.showNext) {
                        url = options.authUrl + "?next=" + +document.location.pathname;
                    } else {
                        url = options.authUrl;
                    }
                    new Snackbar("请<a href='" + url + "'>登录</a>后进行操作", {timeout: options.snackBarAlive});
                }
            }
            this.execute(url, data,options, onPostSuccess, onUnAuth, onPostError, onError,onFinish);
        },
        execute: function (url, data,options,onPostSuccess, onUnAuth, onPostError, onError,onFinish) {
            var xsrf;
            var finish = function(code){
                if (onFinish){
                    onFinish(code);
                }
            };
            try { //cookie may be null or something else bad data
                xsrf = Base64.decode(Cookies.get('_xsrf').split("|")[0]);
            } catch (err) {
                new Snackbar("会话已过期,请<a href='" + window.location.href + "'>刷新</a>重试", {timeout: options.snackBarAlive});
                finish(0);
                return;
            }
            $.ajax({
                type: 'POST',
                url: url,
                data: $.extend({}, {_xsrf: xsrf}, data),
                success: function (data) {
                    try {
                        switch (data.Status) {
                            case 0:
                                onPostError(data.Error);
                                finish(1);
                                break;
                            case 1:
                                if (onPostSuccess) {
                                    onPostSuccess(data);
                                }
                                finish(2);
                                break;
                        }
                    } catch (err) {
                        onError();
                        finish(3);
                    }
                }, error: function (r, err) {
                    if (r.status == 401) {
                        onUnAuth();
                        finish(4);
                    } else {
                        onError();
                        finish(5);
                    }
                }
            });
        }
    },
    simpleParseError: {
        options: {},
        init: function (status, error, options) {
            this.options = $.extend({}, {
                snackAlive: 3000,
                errorCallback: function (message) {
                    new Snackbar(message, {timeout:  this.snackTimeout});
                },
                onSuccess: null
            }, options);
            this.execute(status, error)
        },
        execute: function (status, error) {
            switch (status) {
                case 0:
                    this.options.errorCallback(error);
                    return false;
                case 1:
                    if (this.options.onSuccess != null) {
                        this.options.onSuccess();
                    }
                    return true;
            }
        }
    }
};


function formatTime(value) {
    if (typeof value != "number") {
        var v = Date.parse(value);
        if (isNaN(v)) {
            value = (new Date).getTime();
        } else {
            value = v;
        }
    }
    now = (new Date).getTime();
    if (now - value < 60 * 1000) {
        return "just now";
    }
    if (now - value < 60 * 60 * 1000) {
        var min = parseInt((now - value) / (60 * 1000));
        return min + "minutes ago";
    }
    if (now - value < 24 * 60 * 60 * 1000) {
        var hour = parseInt((now - value) / (60 * 60 * 1000));
        return hour + "hours ago";
    }
    if (now - value < 20 * 24 * 60 * 60 * 1000) {
        var day = parseInt((now - value) / (24 * 60 * 60 * 1000));
        return day + "days ago";
    }
    var d = new Date(value);
    return d.getFullYear()+"-"+(d.getMonth()+1)+"-"+d.getDate();
}