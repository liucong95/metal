/*
Navicat MySQL Data Transfer

Source Server         : alimysql
Source Server Version : 50723
Source Host           : 112.74.171.94:3306
Source Database       : issue

Target Server Type    : MYSQL
Target Server Version : 50723
File Encoding         : 65001

Date: 2018-12-26 11:16:34
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `article`
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL DEFAULT '',
  `content` text NOT NULL,
  `status` tinyint(3) unsigned NOT NULL DEFAULT '1',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of article
-- ----------------------------
INSERT INTO `article` VALUES ('1', 'Go语言优势', '<p>一、\n                  部署简单。Go 编译生成的是一个静态可执行文件，除了 glibc 外没有其他外部依赖。这让部署变得异常方便：目标机器上只需要一个基础的系统和必要的管理、\n                  监控工具，完全不需要操心应用所需的各种包、库的依赖关系，大大减轻了维护的负担。这和 Python 有着巨大的区别。由于历史的原因，\n                  Python 的部署工具生态相当混乱【比如 setuptools, distutils, pip, buildout 的不同适用场合以及兼容性问题】。官方 PyPI 源又经常出问题，\n                  需要搭建私有镜像，而维护这个镜像又要花费不少时间和精力。\n              </p>\n              <p>二、并发性好。Goroutine 和 channel 使得编写高并发的服务端软件变得相当容易，很多情况下完全不需要考虑锁机制以及由此带来的各种问题。单个 Go 应用也能有效的利用多个 CPU 核，并行执行的性能好。这和 Python 也是天壤之比。多线程和多进程的服务端程序编写起来并不简单，而且由于全局锁 GIL 的原因，多线程的 Python 程序并不能有效利用多核，只能用多进程的方式部署；如果用标准库里的 multiprocessing 包又会对监控和管理造成不少的挑战【我们用的 supervisor 管理进程，对 fork 支持不好】。部署 Python 应用的时候通常是每个 CPU 核部署一个应用，这会造成不少资源的浪费，比如假设某个 Python 应用启动后需要占用 100MB 内存，而服务器有 32 个 CPU 核，那么留一个核给系统、运行 31 个应用副本就要浪费 3GB 的内存资源。\n              </p>\n              <p>三、良好的语言设计。从学术的角度讲 Go 语言其实非常平庸，不支持许多高级的语言特性；但从工程的角度讲，Go 的设计是非常优秀的：规范足够简单灵活，有其他语言基础的程序员都能迅速上手。更重要的是 Go 自带完善的工具链，大大提高了团队协作的一致性。比如 gofmt 自动排版 Go 代码，很大程度上杜绝了不同人写的代码排版风格不一致的问题。把编辑器配置成在编辑存档的时候自动运行 gofmt，这样在编写代码的时候可以随意摆放位置，存档的时候自动变成正确排版的代码。此外还有 gofix, govet 等非常有用的工具。\n              </p>\n              <p>四、执行性能好。虽然不如 C 和 Java，但通常比原生 Python 应用还是高一个数量级的，适合编写一些瓶颈业务。内存占用也非常省。</p>', '1', '2018-10-24 11:42:33', '2018-10-24 11:42:34');
INSERT INTO `article` VALUES ('2', '各平台mysql重启', 'linux平台及windows平台mysql重启方法\r\n\r\n　　Linux下重启MySQL的正确方法：\r\n\r\n　　1、通过rpm包安装的MySQL\r\n\r\n　　service mysqld restart\r\n\r\n　　2、从源码包安装的MySQL\r\n\r\n　　// linux关闭MySQL的命令\r\n\r\n　　$mysql_dir/bin/mysqladmin -uroot -p shutdown\r\n\r\n　　// linux启动MySQL的命令\r\n\r\n　　$mysql_dir/bin/mysqld_safe &\r\n\r\n　　其中mysql_dir为MySQL的安装目录，mysqladmin和mysqld_safe位于MySQL安装目录的bin目录下，很容易找到的。\r\n\r\n　　3、以上方法都无效的时候，可以通过强行命令：“killall mysql”来关闭MySQL，但是不建议用这样的方式，因为这种野蛮的方法会强行终止MySQL数据库服务，有可能导致表损坏\r\n\r\n　　步骤或方法:RedHat Linux (Fedora Core/Cent OS)\r\n\r\n　　1.启动：/etc/init.d/mysqld start\r\n\r\n　　2.停止：/etc/init.d/mysqld stop\r\n\r\n　　3.重启：/etc/init.d/mysqld restart\r\n\r\n　　Debian / Ubuntu Linux\r\n\r\n　　1.启动：/etc/init.d/mysql start\r\n\r\n　　2.停止：/etc/init.d/mysql stop\r\n\r\n　　3.重启：/etc/init.d/mysql restart\r\n\r\n　　Windows\r\n\r\n　　1.点击“开始”->“运行”(快捷键Win+R)。\r\n\r\n　　2.启动：输入 net stop mysql\r\n\r\n　　3.停止：输入 net start mysql\r\n\r\n　　提示* Redhat Linux 也支持service command，启动：# service mysqld start 停止：# service mysqld stop 重启：# service mysqld restart\r\n\r\n　　* Windows下不能直接重启(restart)，只能先停止，再启动。\r\n\r\n　　MySQL启动，停止，重启方法：\r\n\r\n　　一、启动方式\r\n\r\n　　1、使用 service 启动：service mysqld start\r\n\r\n　　2、使用 mysqld 脚本启动：/etc/inint.d/mysqld start\r\n\r\n　　3、使用 safe_mysqld 启动：safe_mysqld&\r\n\r\n　　二、停止\r\n\r\n　　1、使用 service 启动：service mysqld stop\r\n\r\n　　2、使用 mysqld 脚本启动：/etc/inint.d/mysqld stop\r\n\r\n　　3、mysqladmin shutdown\r\n\r\n　　三、重启\r\n\r\n　　1、使用 service 启动：service mysqld restart\r\n\r\n　　2、使用 mysqld 脚本启动：/etc/inint.d/mysqld restart', '1', '2018-10-11 16:47:02', '2018-10-11 16:47:02');
INSERT INTO `article` VALUES ('6', '说道说道前后端分离 ', '<p>\n<p>要说前端界的发展速度，那真是快！\n2012年那时候接触过extjs，用于企业级后台开发还真不错，有好看的UI界面，组件丰富，基本能满足各类需求。但此时，HTML5正在蓬勃发展，尤其是乔布斯宣布苹果设备不支持flash后HTML5发展更是迅猛。并且angularjs这类MVVM框架被大多数所知，reactjs,vuejs如雨后春笋般生长。</p>\n<p>2014年使用了一段时间angularjs，感觉学习难度有点大，2015年使用vue1.0做了一个项目后我逢人就说angular，vue有多好用，推荐他们放弃jquery使用vue。不到2年时间再看看前端界，vue，react等框架已经是前端开发标配，如果你说公司项目还在使用jquery会被人笑话，对于前端新人MVVM框架是必学，jquery反而不会重视。</p>\n<p>这就导致一个结果：“手里拿着锤子，看什么都是钉子”，因为有锤子的关系，遇到任何问题，都会先想如何用锤子解决。久而久之，陷入了一种思维定式。任何工具带来便利的同时，也带来了局限性。而这往往是用锤子的人很难看到的。就拿一个需要SEO的网站来说，该选择哪种技术好？如果只会vue，不熟悉jquery的人来说肯定是选择vue，就算vue不太适合做这类网站，也会拿出ssr来强行做事。殊不知需要SEO的网站使用静态文件是最合适的。大多数人认为前后端分离是使用vue，react，angular，使用jquery的不叫前后端分离，这完全是搞混了概念，实际上后端的controller层也属于视图层，也可以归属于前端。</p>\n<p>现在去网上一搜，问一下身边的人怎么看待前后端分离的，大多数人秉持着支持的态度，认为前后端分离好处多多，列举几条：</p>\n<ul>\n<li>专业的人做专业的事</li>\n<li>前后并行开发，效率高</li>\n<li>前端工程化，组件化</li>\n<li>解耦</li>\n<li>降低了开发学习难度</li>\n</ul>\n<p>等等……\n<img alt=\"\" src=\"http://hopefully-img.yuedun.wang/%E5%BE%AE%E4%BF%A1%E5%9B%BE%E7%89%87_20180805104325.jpg\">\n大家说的都说到点上了，这也是前后端分离能发展起来的驱动力。但是道理说的都挺好，如果不结合实际情况的话就是大炮打蚊子，不但达不到理想的效果还浪费资源。而且前后端分离带来的负面情况也不可忽视：</p>\n<ul>\n<li>增加了沟通成本，一些前后端都可以做又都不想做的事或许只能由权利大的来决定。</li>\n<li>一般前端开发速度会比后端快，在接口没开发好时前端只能闲等着。也有反过来的情况。前后端分离也意味着任务关联性减弱，可能不是同时开发，需要一方催着一方来完成。</li>\n<li>跨域问题导致联调困难，前端只能等待接口开发上了服务器才能调试。</li>\n<li>职责分离后确认职责也困难，一个问题出现到底是谁的问题？谁解决？</li>\n<li>一个需求需要前后端开发同时参与理解需求，有理解偏差问题。</li>\n<li>小公司多一个人多一份支出。</li>\n</ul>\n<p>那么到底该不该进行前后端分离，如何进行技术选型？这需要根据一些实际情况来决定，大体判断准则有以下几点：</p>\n<ul>\n<li>后台系统采用前后端分离比较合适。</li>\n<li>需要SEO引流的就不要强行前后端分离了，react，vue的服务端渲染也很勉强，徒增开发难度而已。</li>\n<li>数据交互比较多的使用前后端分离，操作数据比jquery方便。</li>\n<li>页面本身特别简单，只负责简单数据展示，要求打开速度，直接服务端渲染即可。这种页面本身就是单页面，如果还要使用框架就是多此一举，增加页面负担，增加开发调试难度。</li>\n<li>开发资源充足最好前后端分离，开发资源不足时不分离，一人包揽前后台端反而更快。</li>\n</ul>\n<p>最后，前后端分离是一个趋势，但不是必须。更准确的说法应该叫做“前后端分工”，毕竟在5年前这些活都是一个开发来做的，因为技术复杂性提升，前端不想只是切图，后端不想学变化太快的前端就出现了分离。你可以想象测试的工作，现在的测试大多还是测业务，但是也出现了一个自动化测试的职位，因为测试不想天天鼠标点呀点的测，想搞点高深的东西，而开发又特别烦写单元测试代码，这就又出现分离。再者，数据库也是一样，所以出现了DBA这个角色。谁知哪一天又会合起来呢！</p>\n\n<br></p>', '1', '2018-10-12 12:19:51', '2018-10-12 12:19:51');

-- ----------------------------
-- Table structure for `assistance`
-- ----------------------------
DROP TABLE IF EXISTS `assistance`;
CREATE TABLE `assistance` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) DEFAULT NULL,
  `description` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `user_agent` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `referer` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `urgency_level` int(11) NOT NULL DEFAULT '1',
  `state` int(11) NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `images` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `first_helper` int(11) DEFAULT '0',
  `second_helper` int(11) DEFAULT '0',
  `user_account` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Records of assistance
-- ----------------------------
INSERT INTO `assistance` VALUES ('1', '1', '从页面显示效果来看，被 <b> 和 <strong> 包围的文字将会被加粗，而被 <i> 和 <em> 包围的文字将以斜体的形式呈现。那大家可能就会疑惑了，既然效果一样，那为什么还要重复定义标签呢？这就要从 HTML5 的一个最大的特性 -- 语义化来谈了。\n\n作者：编译青春\n链接：https://www.zhihu.com/question/19551271/answer/146298923\n来源：知乎\n著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。', null, null, '1', '0', '2017-10-21 04:06:00', '2017-10-21 04:06:00', null, '0', '0', null);
INSERT INTO `assistance` VALUES ('2', '1', '不得不说，<b> 和 <i> 创建之初就是简单地表示粗体和斜体样式，但现在是 HTML5 的天下。语义化是 HTML5 最大的特性之一，而所有被 HTML5 保留的标签都带有其特有的语义，<b> 和 <i> 也不例外，它们分别被重新赋予了语义。相比较而言，标签的样式反而变得无足轻重，所以上面所讲的两组标签，虽然样式上表现极其相似，但其实语义上各有侧重。\n\n作者：编译青春\n链接：https://www.zhihu.com/question/19551271/answer/146298923\n来源', null, null, '1', '0', '2017-10-21 04:14:30', '2017-10-21 04:14:30', null, '0', '0', null);
INSERT INTO `assistance` VALUES ('6', '1', 'XSS攻击全称跨站脚本攻击，是为不和层叠样式表(Cascading Style Sheets, CSS)的缩写混淆， 故将跨站脚本攻击缩写为XSS，XSS是一种在web应用中的计算机安全漏洞，它允许恶意web用户将代码植入到 提供给其它用户使用的页面中。从而达到攻击的目的。如，盗取用户Cookie、破坏页面结构、重定向到其它网站等。 XSS攻击案例： 新浪微博遭受XX攻击', null, 'http://localhost:3002/', '1', '0', '2017-11-18 08:32:58', '2017-11-18 08:32:58', null, '0', '0', null);
INSERT INTO `assistance` VALUES ('7', '1', 'XSS攻击全称跨站脚本攻击，是为不和层叠样式表(Cascading Style Sheets, CSS)的缩写混淆， 故将跨站脚本攻击缩写为XSS，XSS是一种在web应用中的计算机安全漏洞，它允许恶意web用户将代码植入到 提供给其它用户使用的页面中。从而达到攻击的目的。如，盗取用户Cookie、破坏页面结构、重定向到其它网站等。 XSS攻击案例： 新浪微博遭受XX攻击', null, 'http://localhost:3002/', '1', '0', '2017-11-18 08:33:26', '2017-11-18 08:33:26', null, '0', '0', null);
INSERT INTO `assistance` VALUES ('8', '1', 'http://soft.yesky.com/security/156/30179156.shtml 人人网遭受XSS攻击： http://www.freebuf.com/articles/6295.html 简单的测试方法： 所有提交数据的地方都有可能存在XSS，', null, 'http://localhost:3002/', '1', '0', '2017-11-18 08:34:15', '2017-11-18 08:34:15', null, '0', '0', null);
INSERT INTO `assistance` VALUES ('9', '1', 'http://soft.yesky.com/security/156/30179156.shtml 人人网遭受XSS攻击： http://www.freebuf.com/articles/6295.html 简单的测试方法： 所有提交数据的地方都有可能存在XSS，', null, 'http://localhost:3002/', '1', '0', '2017-11-18 08:39:30', '2017-11-18 08:39:30', null, '0', '0', null);
INSERT INTO `assistance` VALUES ('10', '1', '诱导或等待用户去点击链接，才能触发XSS代码，达到劫持访问、获取cookies的目的。一般容易出现在搜索页面。 例如：https://m.wuage.com/search/self-shop?memberId=4lv8ll4g&keywords=x%22alert(1)%22&psa=M3.s10.0.j4 （此漏洞已经修复，请勿再测，造成访问压力。） 图片描述 ', null, 'http://localhost:3002/', '1', '0', '2017-11-18 08:52:44', '2017-11-18 08:52:44', null, '0', '0', null);
INSERT INTO `assistance` VALUES ('11', '1', '库或者文件中，成为某个url正常的页面的一部分，所有访问这个页面的所有用户都是受害者，看似正常的url，则其页面已经包含了xss代码，持久型XSS更具有隐蔽性，带来的危害也更大 例如：在页面中不容注意的地方加一段js脚本（如下），当页面被打开时，页面会加载这段脚本，加系统登录', null, 'http://localhost:3002/platform/assistance-list', '1', '0', '2017-11-19 06:28:57', '2017-11-19 06:28:57', null, '0', '0', null);
INSERT INTO `assistance` VALUES ('12', '1', 'script节点的语句，载入来自第三方域的含有 具体恶意代码的脚本。具体的恶意代码，常见的行为是读取cookie，构造例如一个img标签，将其src属性指向 恶意第三方网站，将cookie的内容作为参数附在src的url上，这样黑客就能在其网', null, 'http://localhost:3002/platform/assistance-list', '1', '0', '2017-11-19 06:31:23', '2017-11-19 06:31:23', null, '0', '0', null);
INSERT INTO `assistance` VALUES ('13', '1', 'XSS攻击： http://www.freebuf.com/articles/6295.html 简单的测试方法： 所有提交数据的地方都有可能存在XSS，可以用最简单脚本进行测试： XSS攻击类型 反射型： 黑客构造一个包含XSS代码的URL（服务器中没', null, 'http://localhost:3002/', '1', '0', '2017-11-19 10:15:26', '2017-11-19 10:15:26', null, '0', '0', null);
INSERT INTO `assistance` VALUES ('14', '1', '此漏洞已经修复，请勿再测，造成访问压力。） 图片描述 持久型： 如果黑客可以将脚本代码通过发布内容（如发论坛、博文、写留言等）的方式发布后，存储在服务端的数据库或者文件中，成为某个url正常的页面的一部分，所有访问这个页面的所有用户都是受害者，看似正常的url，则其页面已经包含了xss代码，持久型XSS更具有隐蔽性，带来的危', null, 'http://localhost:3002/platform/assistance-list', '1', '0', '2017-11-19 10:17:02', '2017-11-19 10:17:02', null, '0', '0', null);
INSERT INTO `assistance` VALUES ('15', '1', 'dom xss并不复杂，他也属于反射型xss的一种(，domxss取决于输出位置，并不取决于输出环境， 因此domxss既有可能是反射型的，也有可能是存储型的)，简单去理解就是因为他输出点在DOM，所以', null, 'http://localhost:3002/platform/assistance-list', '1', '0', '2018-03-03 15:58:10', '2018-03-03 15:58:10', null, '1', '0', null);

-- ----------------------------
-- Table structure for `assistance_people`
-- ----------------------------
DROP TABLE IF EXISTS `assistance_people`;
CREATE TABLE `assistance_people` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_name` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `mobile` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `email` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `superior` int(11) DEFAULT NULL,
  `features` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Records of assistance_people
-- ----------------------------

-- ----------------------------
-- Table structure for `features`
-- ----------------------------
DROP TABLE IF EXISTS `features`;
CREATE TABLE `features` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `feature_name` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Records of features
-- ----------------------------
INSERT INTO `features` VALUES ('1', '教学工作站', '2017-11-18 15:28:10', '2017-11-18 15:28:13');
INSERT INTO `features` VALUES ('2', '销售工作站', '2017-11-18 15:28:25', '2017-11-18 15:28:30');
INSERT INTO `features` VALUES ('3', '市场工作站', '2017-11-18 15:28:42', '2017-11-18 15:28:45');
INSERT INTO `features` VALUES ('4', '数据管理', '2017-11-18 15:28:54', '2017-11-18 15:28:57');
INSERT INTO `features` VALUES ('5', '管理中心', '2017-11-18 15:29:08', '2017-11-18 15:29:11');
INSERT INTO `features` VALUES ('6', '用户管理', '2017-11-18 15:29:19', '2017-11-18 15:29:22');
INSERT INTO `features` VALUES ('7', '财务工作站', '2017-11-18 15:31:06', '2017-11-18 15:31:09');

-- ----------------------------
-- Table structure for `helper_feature_relation`
-- ----------------------------
DROP TABLE IF EXISTS `helper_feature_relation`;
CREATE TABLE `helper_feature_relation` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `helper_id` int(11) DEFAULT NULL,
  `feature_id` int(11) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Records of helper_feature_relation
-- ----------------------------

-- ----------------------------
-- Table structure for `helpers`
-- ----------------------------
DROP TABLE IF EXISTS `helpers`;
CREATE TABLE `helpers` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_name` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `mobile` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `email` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `superior` int(11) DEFAULT NULL,
  `features` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Records of helpers
-- ----------------------------
INSERT INTO `helpers` VALUES ('1', '张三', '18734587454', 'zhangsan@qq.com', null, '教学工作站,销售工作站', '2017-10-24 23:11:10', '2017-10-24 23:11:14');
INSERT INTO `helpers` VALUES ('2', '李四', '12222222222', '12222222222@163.com', null, '财务工作站,用户管理', '2017-11-18 15:30:06', '2017-11-18 15:30:09');

-- ----------------------------
-- Table structure for `job_count`
-- ----------------------------
DROP TABLE IF EXISTS `job_count`;
CREATE TABLE `job_count` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `job_title` varchar(30) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '职位名称，开发语言',
  `region` varchar(10) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '地区',
  `amount` int(11) NOT NULL DEFAULT '0' COMMENT '职位数',
  `created_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=323 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Records of job_count
-- ----------------------------
INSERT INTO `job_count` VALUES ('1', 'nodejs', '上海', '120', '2018-03-13 11:30:21');
INSERT INTO `job_count` VALUES ('2', 'golang', '上海', '90', '2018-03-13 11:30:44');
INSERT INTO `job_count` VALUES ('3', 'nodejs', '上海', '130', '2018-03-25 11:31:14');
INSERT INTO `job_count` VALUES ('4', 'golang', '上海', '110', '2018-03-25 11:31:43');
INSERT INTO `job_count` VALUES ('5', 'nodejs', '上海', '141', '2018-03-30 11:32:14');
INSERT INTO `job_count` VALUES ('6', 'golang', '上海', '117', '2018-03-30 11:32:40');
INSERT INTO `job_count` VALUES ('7', 'nodejs', '上海', '138', '2018-04-02 11:33:21');
INSERT INTO `job_count` VALUES ('8', 'golang', '上海', '116', '2018-04-02 11:33:44');
INSERT INTO `job_count` VALUES ('9', 'nodejs', '上海', '147', '2018-04-05 11:34:57');
INSERT INTO `job_count` VALUES ('10', 'golang', '上海', '128', '2018-04-05 11:35:17');
INSERT INTO `job_count` VALUES ('11', 'nodejs', '上海', '141', '2018-04-10 11:35:45');
INSERT INTO `job_count` VALUES ('12', 'golang', '上海', '130', '2018-04-10 11:36:04');
INSERT INTO `job_count` VALUES ('13', 'nodejs', '上海', '139', '2018-04-14 11:36:23');
INSERT INTO `job_count` VALUES ('14', 'golang', '上海', '133', '2018-04-14 11:36:40');
INSERT INTO `job_count` VALUES ('15', 'nodejs', '上海', '153', '2018-04-23 11:37:23');
INSERT INTO `job_count` VALUES ('16', 'golang', '上海', '143', '2018-04-23 11:37:44');
INSERT INTO `job_count` VALUES ('17', 'nodejs', '上海', '155', '2018-04-27 11:38:02');
INSERT INTO `job_count` VALUES ('18', 'golang', '上海', '144', '2018-04-27 11:38:26');
INSERT INTO `job_count` VALUES ('19', 'nodejs', '上海', '148', '2018-05-26 11:38:53');
INSERT INTO `job_count` VALUES ('20', 'golang', '上海', '165', '2018-05-26 11:39:24');
INSERT INTO `job_count` VALUES ('21', 'golang', '上海', '163', '2018-06-21 12:00:01');
INSERT INTO `job_count` VALUES ('22', 'nodejs', '上海', '135', '2018-06-21 12:00:01');
INSERT INTO `job_count` VALUES ('23', 'golang', '上海', '164', '2018-06-22 12:00:01');
INSERT INTO `job_count` VALUES ('24', 'nodejs', '上海', '137', '2018-06-22 12:00:01');
INSERT INTO `job_count` VALUES ('25', 'golang', '上海', '164', '2018-06-23 12:00:01');
INSERT INTO `job_count` VALUES ('26', 'nodejs', '上海', '138', '2018-06-23 12:00:01');
INSERT INTO `job_count` VALUES ('27', 'nodejs', '上海', '139', '2018-06-24 12:00:01');
INSERT INTO `job_count` VALUES ('28', 'golang', '上海', '164', '2018-06-24 12:00:01');
INSERT INTO `job_count` VALUES ('29', 'golang', '上海', '160', '2018-06-25 12:00:01');
INSERT INTO `job_count` VALUES ('30', 'nodejs', '上海', '138', '2018-06-25 12:00:01');
INSERT INTO `job_count` VALUES ('31', 'golang', '上海', '163', '2018-06-26 12:00:01');
INSERT INTO `job_count` VALUES ('32', 'nodejs', '上海', '138', '2018-06-26 12:00:01');
INSERT INTO `job_count` VALUES ('33', 'nodejs', '上海', '137', '2018-06-27 12:00:01');
INSERT INTO `job_count` VALUES ('34', 'golang', '上海', '163', '2018-06-27 12:00:01');
INSERT INTO `job_count` VALUES ('35', 'golang', '上海', '163', '2018-06-28 12:00:01');
INSERT INTO `job_count` VALUES ('36', 'nodejs', '上海', '135', '2018-06-28 12:00:01');
INSERT INTO `job_count` VALUES ('37', 'golang', '上海', '163', '2018-06-29 12:00:01');
INSERT INTO `job_count` VALUES ('38', 'nodejs', '上海', '138', '2018-06-29 12:00:01');
INSERT INTO `job_count` VALUES ('39', 'nodejs', '上海', '133', '2018-06-30 12:00:01');
INSERT INTO `job_count` VALUES ('40', 'golang', '上海', '166', '2018-06-30 12:00:01');
INSERT INTO `job_count` VALUES ('41', 'nodejs', '上海', '130', '2018-07-01 12:00:01');
INSERT INTO `job_count` VALUES ('42', 'golang', '上海', '166', '2018-07-01 12:00:01');
INSERT INTO `job_count` VALUES ('43', 'nodejs', '上海', '124', '2018-07-02 12:00:01');
INSERT INTO `job_count` VALUES ('44', 'golang', '上海', '164', '2018-07-02 12:00:01');
INSERT INTO `job_count` VALUES ('45', 'nodejs', '上海', '130', '2018-07-03 12:00:01');
INSERT INTO `job_count` VALUES ('46', 'golang', '上海', '167', '2018-07-03 12:00:01');
INSERT INTO `job_count` VALUES ('47', 'nodejs', '上海', '133', '2018-07-04 12:00:01');
INSERT INTO `job_count` VALUES ('48', 'golang', '上海', '168', '2018-07-04 12:00:01');
INSERT INTO `job_count` VALUES ('49', 'nodejs', '上海', '135', '2018-07-05 12:00:01');
INSERT INTO `job_count` VALUES ('50', 'golang', '上海', '169', '2018-07-05 12:00:01');
INSERT INTO `job_count` VALUES ('51', 'nodejs', '上海', '134', '2018-07-07 12:00:01');
INSERT INTO `job_count` VALUES ('52', 'golang', '上海', '173', '2018-07-07 12:00:01');
INSERT INTO `job_count` VALUES ('53', 'golang', '上海', '173', '2018-07-08 12:00:01');
INSERT INTO `job_count` VALUES ('54', 'nodejs', '上海', '133', '2018-07-08 12:00:01');
INSERT INTO `job_count` VALUES ('55', 'nodejs', '上海', '126', '2018-07-09 12:00:01');
INSERT INTO `job_count` VALUES ('56', 'golang', '上海', '168', '2018-07-09 12:00:01');
INSERT INTO `job_count` VALUES ('57', 'nodejs', '上海', '125', '2018-07-10 12:00:01');
INSERT INTO `job_count` VALUES ('58', 'golang', '上海', '172', '2018-07-10 12:00:01');
INSERT INTO `job_count` VALUES ('59', 'golang', '上海', '174', '2018-07-11 12:00:01');

-- ----------------------------
-- Table structure for `log`
-- ----------------------------
DROP TABLE IF EXISTS `log`;
CREATE TABLE `log` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `mark` varchar(255) DEFAULT NULL COMMENT '日志说明',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=47 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of log
-- ----------------------------
INSERT INTO `log` VALUES ('1', '登录IP:43.227.254.20，物理地址：中国  北京 北京', '2018-09-21 00:41:46', '2018-09-21 00:41:46');
INSERT INTO `log` VALUES ('2', '登录IP:101.81.62.69，物理地址：中国  上海 上海', '2018-09-25 14:10:11', '2018-09-25 14:10:11');
INSERT INTO `log` VALUES ('3', '登录IP:116.226.177.39，物理地址：中国  上海 上海', '2018-09-28 09:38:48', '2018-09-28 09:38:48');
INSERT INTO `log` VALUES ('4', '登录IP:222.209.182.112，物理地址：中国  四川 成都', '2018-09-29 17:44:54', '2018-09-29 17:44:54');

-- ----------------------------
-- Table structure for `role`
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `privileges` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '权限列表',
  `description` varchar(255) COLLATE utf8_unicode_ci NOT NULL COMMENT '说明',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='可以给一个角色很多权限，也可以通过很多角色组合来拥有很多权限';

-- ----------------------------
-- Records of role
-- ----------------------------
INSERT INTO `role` VALUES ('1', '*', '管理员', '2018-04-21 14:00:18', '2018-04-21 14:00:23');
INSERT INTO `role` VALUES ('2', 'ArticleController:*,OfficialController:*', '空间业务查看/发表', '2018-04-21 14:00:18', '2018-04-21 14:00:23');
INSERT INTO `role` VALUES ('3', 'ArticleController:ArticlesRoute,ArticleController:ArticlesList', '空间业务查看', '2018-04-21 14:00:15', '2018-04-21 14:00:21');

-- ----------------------------
-- Table structure for `user`
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `account` varchar(255) NOT NULL DEFAULT '',
  `password` varchar(255) NOT NULL DEFAULT '',
  `user_name` varchar(255) NOT NULL DEFAULT '',
  `mobile` varchar(255) NOT NULL DEFAULT '',
  `email` varchar(255) NOT NULL DEFAULT '',
  `status` int(11) NOT NULL DEFAULT '1',
  `description` varchar(255) NOT NULL DEFAULT '',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_account_uniq` (`account`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('1', 'admin', '21232f297a57a5a743894a0e4a801fc3', 'admin', '110', '13477889900@139.com', '1', '管理员', '2018-03-17 20:46:31', '2018-11-01 01:31:17');
INSERT INTO `user` VALUES ('2', 'user1', '21232f297a57a5a743894a0e4a801fc3', 'joker', '110', '18701897513@139.com', '1', 'xxxxx', '2018-03-17 20:49:44', '2018-06-11 12:56:34');
INSERT INTO `user` VALUES ('3', 'user2', '21232f297a57a5a743894a0e4a801fc3', 'huni', '110', '18611118146@139.com', '1', '', '2017-07-27 03:25:01', '2018-02-28 11:57:56');
INSERT INTO `user` VALUES ('4', '10701897527', '21232f297a57a5a743894a0e4a801fc3', '缇欧', '110', 'huo@gmail.com', '1', 'hello world', '2017-07-27 09:00:43', '2018-03-19 11:10:50');
-- ----------------------------
-- Table structure for `privilege`
-- ----------------------------
DROP TABLE IF EXISTS `privilege`;
CREATE TABLE `privilege` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `role_id` int(11) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_privilege` (`user_id`,`role_id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Records of privilege
-- ----------------------------
INSERT INTO `privilege` VALUES ('1', '1', '1', '2018-04-28 17:35:30', '2018-04-28 17:35:32');
INSERT INTO `privilege` VALUES ('2', '2', '2', '2018-04-28 13:16:33', '2018-04-28 13:16:36');
INSERT INTO `privilege` VALUES ('3', '3', '3', '2018-06-11 13:10:12', '2018-06-11 13:10:15');

