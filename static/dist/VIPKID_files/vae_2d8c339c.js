webpackJsonp([18], {
    1: function (e, t, a) {
        "use strict";
        var i = a(0),
            o = function (e) {
                return e && e.__esModule ? e : {
                    default:
                        e
                }
            }(i); (0, o.
                default)(function () {
                    ({
                        isIpadBox:
                            !0,
                        init: function () {
                            this.eventClick(),
                                this.skip(),
                                this.eventClickToggleNavigation(),
                                this.padMenu(),
                                this.registerJump(),
                                this.loginJump()
                        },
                        registerJump: function () {
                            (0, o.
                                default)("#registerHeaderBtn,#padRegisterHeaderBtn").on("click",
                                    function () {
                                        vkTrack.click("parent_pc_signup_click_head"),
                                            window.location.href = "/signup"
                                    })
                        },
                        loginJump: function () {
                            (0, o.
                                default)("#loginHeaderBtn,#padLoginHeaderBtn").on("click",
                                    function () {
                                        vkTrack.click("parent_pc_login_click_head"),
                                            window.location.href = "/login"
                                    })
                        },
                        padMenu: function () {
                            (0, o.
                                default)("#systemMenu").click(function () {
                                    (0, o.
                                        default)("#systemMenu .menu-box").toggle()
                                })
                        },
                        eventClick: function () {
                            var e = !0; (0, o.
                                default)("#navBtn").on("click",
                                    function () {
                                        e ? ((0, o.
                                            default)("body,html").css("overflow", "hidden"), (0, o.
                                                default)("body").addClass("active"), (0, o.
                                                    default)("#navLiderLeft").addClass("sliderbarShow"), e = !1) : ((0, o.
                                                        default)("body,html").css("overflow", "auto"), (0, o.
                                                            default)("body").removeClass("active"), (0, o.
                                                                default)("#navLiderLeft").removeClass("sliderbarShow"), e = !0)
                                    })
                        },
                        getRequest: function () {
                            var e = window.location.search,
                                t = new Object;
                            if (- 1 != e.indexOf("?")) for (var a = e.substr(1), i = a.split("&"), o = 0; o < i.length; o++) t[i[o].split("=")[0]] = decodeURI(i[o].split("=")[1]);
                            return t
                        },
                        skip: function () {
                            var e = {};
                            if (window.location.search.substr(1).split("&").forEach(function (t) {
                                var a = t.split("=");
                                e[a[0]] = a[1]
                            }), e.hasOwnProperty("channel_id")) {
                                var t = e.channel_id; (0, o.
                                    default)(".header-info-pc a, .header-info-ipad a").each(function () {
                                        var e = (0, o.
                                            default)(this).attr("href"); (0, o.
                                                default)(this).attr("href", e + "?channel_id=" + t)
                                    })
                            }
                        },
                        eventClickToggleNavigation: function () {
                            var e = window.location.pathname;
                            if (e.length <= 1) return void (0, o.
                                default)(".sy").addClass("active");
                            switch (e = e.split("/")[2]) {
                                case "home":
                                    (0, o.
                                        default)(".sy").addClass("active");
                                    break;
                                case "teachers":
                                    (0, o.
                                        default)(".bmsz").addClass("active");
                                    break;
                                case "advantage":
                                    (0, o.
                                        default)(".kctx").addClass("active"),
                                        (0, o.
                                            default)(".menu-box a").css("font-weight", "400"),
                                        (0, o.
                                            default)(".kctxys").css({
                                                color:
                                                    "#F85415",
                                                "font-weight": "600"
                                            });
                                    break;
                                case "mc":
                                    (0, o.
                                        default)(".kctx").addClass("active"),
                                        (0, o.
                                            default)(".menu-box a").css("font-weight", "400"),
                                        (0, o.
                                            default)(".zxk").css({
                                                color:
                                                    "#F85415",
                                                "font-weight": "600"
                                            });
                                    break;
                                case "vae":
                                    (0, o.
                                        default)(".kctx").addClass("active"),
                                        (0, o.
                                            default)(".menu-box a").css("font-weight", "400"),
                                        (0, o.
                                            default)(".qxjjkc").css({
                                                color:
                                                    "#F85415",
                                                "font-weight": "600"
                                            });
                                    break;
                                case "classlist":
                                    (0, o.
                                        default)(".gkk").addClass("active");
                                    break;
                                case "step":
                                    (0, o.
                                        default)(".rhsk").addClass("active");
                                    break;
                                case "download":
                                    (0, o.
                                        default)(".xzzx").addClass("active");
                                    break;
                                case "aboutus":
                                case "news":
                                    (0, o.
                                        default)(".zxdt").addClass("active")
                            }
                        }
                    }).init()
                })
    },
    4: function (e, t, a) {
        "use strict";
        var i = a(0),
            o = function (e) {
                return e && e.__esModule ? e : {
                    default:
                        e
                }
            }(i); (0, o.
                default)(function () {
                    var e = !1; ({
                        init: function () {
                            this.scrollScreenHeight(),
                                this.httpClickRegisterAuditions(),
                                this.goBackTop(),
                                this.eventClickToggleRightFix(),
                                this.eventClickCloseBottom()
                        },
                        goBackTop: function () {
                            (0, o.
                                default)("#goBackBtn").click(function () {
                                    return (0, o.
                                        default)("html,body").animate({
                                            scrollTop:
                                                0
                                        },
                                            1e3),
                                        !1
                                })
                        },
                        scrollScreenHeight: function () {
                            (0, o.
                                default)(window).scroll(function () {
                                    (0, o.
                                        default)(window).scrollTop() >= .5 * (0, o.
                                            default)(window).height() && !e ? ((0, o.
                                                default)("#sectionBottom").addClass("active"), (0, o.
                                                    default)("#sectionBottom").removeClass("hideB"), (0, o.
                                                        default)("#goBackBtn").show()) :
                                    ((0, o.
                                        default)("#sectionBottom").removeClass("active"), (0, o.
                                            default)("#sectionBottom").addClass("hideB"), (0, o.
                                                default)("#goBackBtn").hide()),
                                    (0, o.
                                        default)(window).scrollTop() >= .5 * (0, o.
                                            default)(window).height() ? (0, o.
                                                default)("#goBackBtn").show() :
                                        (0, o.
                                            default)("#goBackBtn").hide(),
                                    (0, o.
                                        default)(window).scrollTop() + (0, o.
                                            default)(window).height() >= (0, o.
                                                default)(document).height() - (0, o.
                                                    default)(".footer").height() && ((0, o.
                                                        default)("#sectionBottom").removeClass("active"), (0, o.
                                                            default)("#sectionBottom").addClass("hideB"), (0, o.
                                                                default)("#goBackBtn").show())
                                })
                        },
                        compile: function (e) {
                            for (var t = String.fromCharCode(e.charCodeAt(0) + e.length), a = 1; a < e.length; a++) t += String.fromCharCode(e.charCodeAt(a) + e.charCodeAt(a - 1));
                            return escape(t)
                        },
                        IsPcPage: function () {
                            var e = 0,
                                t = navigator.userAgent,
                                a = ["Android", "iPhone", "SymbianOS", "Windows Phone", "iPad", "iPod"],
                                i = !0;
                            for (this.isPcBox = !0, e = 0; e < a.length; e++) if (t.indexOf(a[e]) > 0) {
                                i = !1,
                                    this.isPcBox = !1;
                                break
                            }
                            return i
                        },
                        httpClickRegisterAuditions: function () {
                            var e = /^1[3-9][0-9]{9}$/,
                                t = this; (0, o.
                                    default)("#fixClickBottomBtn").click(function () {
                                        var a = (0, o.
                                            default)(".registerAuditions").val();
                                        a && e.test(a) ? (a = t.compile(a), vkTrack.click("parent_pc_bottom_signup_click_box"), location.href = "/signup?&mobile=" + a) : (0, o.
                                            default)("#bottomErrorTips").show()
                                    }),
                                    (0, o.
                                        default)("#fixRightBtn").click(function () {
                                            vkTrack.click("parent_pc_consult_fix_click_th")
                                        })
                        },
                        eventClickToggleRightFix: function () {
                            (0, o.
                                default)("#qrcodeTouchBtn").on("touchend",
                                    function () {
                                        (0, o.
                                            default)("#erweima").toggle()
                                    }),
                            (0, o.
                                default)("#telTouchBtn").on("touchend",
                                    function () {
                                        (0, o.
                                            default)("#teltips").toggle()
                                    })
                        },
                        eventClickCloseBottom: function () {
                            (0, o.
                                default)("#closeBtnBottom").click(function () {
                                    e = !0,
                                        vkTrack.click("parent_pc_bottom_signup_click_close"),
                                        (0, o.
                                            default)("#sectionBottom").removeClass("active"),
                                        (0, o.
                                            default)("#sectionBottom").addClass("hideB")
                                })
                        }
                    }).init()
                })
    },
    6: function (e, t, a) {
        "use strict";
        function i() {
            (0, l.
                default)("body").delegate("[data-video]", "click",
                    function () {
                        function e() {
                            return !!document.createElement("video").canPlayType
                        }
                        function t() {
                            var t = "",
                                a = "";
                            n.hasClass("hide") || (t = (0, l.
                                default)(window).width(), a = (0, l.
                                    default)(window).height(), n.css({
                                        width: t,
                                        height: a,
                                        display: "block"
                                    }), e() ? n.children("video").css({
                                        width: t,
                                        height: a
                                    }) : (n.children("object").attr("style", "width:" + t + "px;height:" + a + "px;"), n.find("embed").attr("style", "width:" + t + "px;height:" + a + "px;")))
                        }
                        var a = "",
                            i = "",
                            o = "",
                            d = "",
                            n = "",
                            s = "",
                            c = "",
                            u = "",
                            r = "",
                            f = "",
                            h = "",
                            v = "",
                            p = ""; (0, l.
                                default)("body").css("overflow", "hidden"),
                                a = (0, l.
                                    default)(window).width(),
                                i = (0, l.
                                    default)(window).height(),
                                o = (0, l.
                                    default)(this).attr("data-video"),
                                d = (0, l.
                                    default)(this).attr("data-video-flv"),
                                n = (0, l.
                                    default)('<div id="videoPop"></div>'),
                                s = (0, l.
                                    default)('<div id="videoCloseBtn"><img style="width: 100%" src="https://image.vipkid.com.cn/market/file/1539676842708-guanbi.png"></div>'),
                                c = {
                                    width: "100%",
                                    position: "fixed",
                                    left: 0,
                                    top: 0,
                                    "z-index": 1010,
                                    display: "none",
                                    background: "rgb(0, 0, 0)",
                                    height: i
                                },
                                u = {
                                    width: 38,
                                    "border-radius": "50%",
                                    position: "absolute",
                                    top: 12,
                                    right: 14,
                                    cursor: "pointer",
                                    transition: "all 0.2s"
                                },
                                r = {
                                    opacity: 1
                                },
                                f = {
                                    color: "#fff",
                                    "text-align": "center",
                                    position: "relative",
                                    "margin-top": "-10px",
                                    top: "50%"
                                },
                                h = {
                                    opacity: .7
                                },
                                n.addClass("hide"),
                                n.css(c),
                                s.css(u),
                                s.hover(function () {
                                    s.css(r)
                                },
                                    function () {
                                        s.css(h)
                                    }),
                                (0, l.
                                    default)("body").append(n),
                                n.removeClass("hide").fadeIn(),
                                e() ? (p = '<video src="' + o + '" type="video/mp4" width="' + a + '" height="' + i + '" autoplay="autoplay" controls="true" style="background:#000 url(\'https://image.vipkid.com.cn/market/file/1539715751879-videoload.gif\') no-repeat center"></video>', n.append(p), n.children("video").bind("ended",
                                    function () {
                                        (0, l.
                                            default)("body").css("overflow", "auto"),
                                        n.removeClass("show").fadeOut(200),
                                        (0, l.
                                            default)(this).parent().addClass("hide").hide(),
                                        (0, l.
                                            default)(this).remove()
                                    })) :
                                    (d ? (v += '<object codebase="//download.macromedia.com/pub/shockwave/cabs/flash/swflash.cab#version=6,0,29,0"  id="videoObject" width="' + a + '" height="' + i + '" classid="clsid:D27CDB6E-AE6D-11cf-96B8-444553540000">', v += '<param name="quality" value="high" />', v += '<param name="allowFullScreen" value="true" />', v += '<param name="wmode" value="transparent" />', v += '<param name="wmode" value="opaque" />', v += '<param name="movie" value="//resource.vipkid.com.cn/parent_portal/images/flvplayer.swf" />', v += '<param name="FlashVars"  value="vcastr_file=' + d + '&amp;IsAutoPlay=1&amp;IsShowBar=0" />', v += '<embed src="//resource.vipkid.com.cn/parent_portal/images/flvplayer.swf" flashvars="vcastr_file=' + d + '&amp;IsAutoPlay=1&amp;IsShowBar=0" width="' + a + '" height="' + i + '" pluginspage=" http://www.macromedia.com/go/getflashplayer" quality="high" allowfullscreen="true" type="application/x-shockwave-flash"  wmode="transparent" /></object>') : (v = (0, l.
                                        default)('<div>暂不支持当前格式的视频播放，请用最新版的<a href="/download" style="color:#fff;margin: 5px">Chrome浏览器</a>打开本站</div>'), v.css(f)), n.append(v)),
                                n.append(s),
                                (0, l.
                                    default)(window).on("resize", t),
                                s.click(function () {
                                    (0, l.
                                        default)("body").css("overflow", "auto"),
                                    (0, l.
                                        default)(this).siblings().remove(),
                                    (0, l.
                                        default)(this).parent().addClass("hide").hide(),
                                    (0, l.
                                        default)(window).unbind("resize", t)
                                })
                    })
        }
        var o = a(0),
            l = function (e) {
                return e && e.__esModule ? e : {
                    default:
                        e
                }
            }(o);
        e.exports = i
    },
    95: function (e, t, a) {
        "use strict";
        function i(e) {
            return e && e.__esModule ? e : {
                default:
                    e
            }
        }
        a(96),
            a(1),
            a(4);
        var o = a(0),
            l = i(o),
            d = a(6),
            n = i(d); (0, l.
                default)(function () {
                    ({
                        init:
                            function () {
                                (0, n.
                                    default)(),
                                this.eventClickVktrackMc(),
                                this.rewardCarousel()
                            },
                        eventClickVktrackMc: function () {
                            (0, l.
                                default)(".freeVipkidMcClassBtn").click(function () {
                                    var e = (0, l.
                                        default)(this).attr("name"),
                                        t = "parent_pc_signup_click_body_box_mc_" + e;
                                    vkTrack.click(t),
                                        window.location.href = "/signup"
                                })
                        },
                        rewardCarousel: function () {
                            var e = (0, l.
                                default)("#swiper-list > li").length,
                                t = (0, l.
                                    default)("#swiper-list > li.active").index(),
                                a = (0, l.
                                    default)(".prve"),
                                i = (0, l.
                                    default)(".next"),
                                o = ["听音选图训练", "连词成句技巧", "语法选择训练", "拼写训练"];
                            a.click(function () {
                                0 === Number(t) ? t = 3 : t -= 1,
                                    (0, l.
                                        default)("#swiper-list li").removeClass("active"),
                                    (0, l.
                                        default)("#swiper-list li").eq(t).addClass("active"),
                                    (0, l.
                                        default)("#focus-bubble li").removeClass("active"),
                                    (0, l.
                                        default)("#focus-bubble li").eq(t).addClass("active"),
                                    (0, l.
                                        default)("#active-text").text(o[t])
                            }),
                                i.click(function () {
                                    Number(t + 1) === e ? t = 0 : t += 1,
                                        (0, l.
                                            default)("#swiper-list li").removeClass("active"),
                                        (0, l.
                                            default)("#swiper-list li").eq(t).addClass("active"),
                                        (0, l.
                                            default)("#focus-bubble li").removeClass("active"),
                                        (0, l.
                                            default)("#focus-bubble li").eq(t).addClass("active"),
                                        (0, l.
                                            default)("#active-text").text(o[t])
                                }),
                                (0, l.
                                    default)("#focus-bubble li").click(function () {
                                        t = (0, l.
                                            default)(this).index(),
                                            (0, l.
                                                default)("#swiper-list li").removeClass("active"),
                                            (0, l.
                                                default)("#swiper-list li").eq(t).addClass("active"),
                                            (0, l.
                                                default)("#focus-bubble li").removeClass("active"),
                                            (0, l.
                                                default)("#focus-bubble li").eq(t).addClass("active"),
                                            (0, l.
                                                default)("#active-text").text(o[t])
                                    })
                        }
                    }).init()
                })
    },
    96: function (e, t) { }
},
    [95]);