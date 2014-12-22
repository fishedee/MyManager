//插件与配置
fis.config.merge({
    modules : {
        postprocessor : {
            js : 'jswrapper'
        },
        postpackager : 'autoload'
    },
    settings : {
        postprocessor : {
            jswrapper : {
                type : 'amd',
            }
        }
    }
});
//压缩打包，减少HTTP请求
fis.config.set('pack', {
	//后台基础库打包
    'pkg/backstageBase.js': [
		'fishstrap/lib/gri/gri.js',
		'fishstrap/lib/gri/griTable.js',
        'fishstrap/core/global.js',
		'fishstrap/core/html5.js',
        'fishstrap/ui/dialog.js',
    ],
	//上传模块打包
	'pkg/upload.js': [
		'fishstrap/util/upload.js',
		'fishstrap/util/jpegEncoder.js',
        'fishstrap/util/imageCompresser.js',
        'fishstrap/util/jpegMeta.js',
    ],
	//ui基础模块打包
	'pkg/uiBase.js': [
		'fishstrap/ui/query.js',
		'fishstrap/ui/input.js',
        'fishstrap/ui/table.js',
        'fishstrap/ui/editor.js',
    ]
});
//目录规范
fis.config.merge({
    roadmap : {
        path : [
            {
                //fishstrap的lib目录下.js文件设置为非模块文件
                reg : /^\/fishstrap\/lib\/(.*)\.js$/i,
                isMod : false,
            },
			{
				//fishstrap的其它目录下js文件设置为模块文件
                reg : /^\/fishstrap\/(.*)\.(js)$/i,
                isMod : true,
			},
			{
                //Makefile文件，不要发布
                reg : /^\/Makefile$/i,
                release : false
            }
        ],
    },
});
//jquery的cdn
fis.config.get('roadmap.path').unshift({
	reg : /^\/fishstrap\/lib\/jquery.js$/i,
	useHash:false,
	isMod:false,
	url:'/libs/jquery/1.11.1/jquery.min.js'
});
fis.config.merge({
	roadmap: {
		domain:{
			'/fishstrap/lib/jquery.js':'http://apps.bdimg.com'
		}
	}
});
//underscore的cdn
fis.config.get('roadmap.path').unshift({
	reg : /^\/fishstrap\/lib\/underscore.js$/i,
	useHash:false,
	isMod:false,
	url:'/libs/underscore.js/1.7.0/underscore-min.js'
});
fis.config.merge({
	roadmap: {
		domain:{
			'/fishstrap/lib/underscore.js':'http://apps.bdimg.com'
		}
	}
});
//echarts的cdn
fis.config.get('roadmap.path').unshift({
	reg : /^\/fishstrap\/lib\/echarts\/echarts.js$/i,
	useHash:false,
	isMod:false,
	url:'/libs/echarts/2.0.4/echarts-plain.js'
});
fis.config.merge({
	roadmap: {
		domain:{
			'/fishstrap/lib/echarts/echarts.js':'http://apps.bdimg.com'
		}
	}
});
//如果要兼容低版本ie显示透明png图片，请使用pngquant作为图片压缩器，
//否则png图片透明部分在ie下会显示灰色背景
//使用spmx release命令时，添加--optimize或-o参数即可生效
//fis.config.set('settings.optimzier.png-compressor.type', 'pngquant');

//设置jshint插件要排除检查的文件，默认不检查lib、jquery、backbone、underscore等文件
//使用spmx release命令时，添加--lint或-l参数即可生效
//fis.config.set('settings.lint.jshint.ignored', [ 'lib/**', /jquery|backbone|underscore/i ]);

//csssprite处理时图片之间的边距，默认是3px
//fis.config.set('settings.spriter.csssprites.margin', 20);
