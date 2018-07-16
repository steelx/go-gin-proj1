var gulp = require('gulp'),
    less = require('gulp-less'),
    livereload = require('gulp-livereload');

gulp.task('less', function() {
  gulp.src('less/app.less')
    .pipe(less())
    .pipe(gulp.dest('public/css'))
    .pipe(livereload());
});

gulp.task('watch', function() {
  livereload.listen({start:true});
  gulp.watch('less/*.less', ['less']);
});
