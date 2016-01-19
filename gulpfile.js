var gulp = require('gulp');

gulp.task('default', function() {
  gulp.src([
    './node_modules/react-bootstrap/dist/react-bootstrap.min.js',
    './node_modules/react/dist/react.min.js',
    './node_modules/react-dom/dist/react-dom.min.js'
  ], {}).pipe(gulp.dest('./build/js'));
  gulp.src([
    './node_modules/bootstrap/dist/css/bootstrap-theme.css',
    './node_modules/bootstrap/dist/css/bootstrap.min.css',
    './node_modules/bootstrap/dist/css/bootstrap-theme.css.map',
    './node_modules/bootstrap/dist/css/bootstrap.min.css.map',
    './src/css/main.css'
  ], {}).pipe(gulp.dest('./build/css'));
});
