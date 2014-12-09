module.exports = function(grunt){
	var now = new Date().valueOf();
	grunt.initConfig({
		pkg:grunt.file.readJSON('package.json'),
		clean:{
			 build:{
				src:['build']
			 }
		},
		copy:{
			build:{
				files:[{expand: true, cwd: 'src/', src: ['**'], dest: 'build/'}]
			}
		},
		replace:{
			build:{
				src:['build/*.html','build/view/**/*.html'],
				overwrite:true,
				replacements:[{
					from:'<?php echo time();?>',
					to: now
				}]
			}
		},
		watch:{
			build:{
				files:'src/**',
				tasks:['default'],
				options:{
					spawn:false,
					livereload:true,
				}
			}
		},
	});
	grunt.loadNpmTasks('grunt-contrib-watch');
	grunt.loadNpmTasks('grunt-contrib-clean');
	grunt.loadNpmTasks('grunt-contrib-copy');
	grunt.loadNpmTasks('grunt-text-replace');
	grunt.registerTask('default',['clean','copy']);
	grunt.registerTask('live',['watch']);
};
