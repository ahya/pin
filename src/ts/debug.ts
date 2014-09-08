///<reference path="../../typings/tsd.d.ts"/>

$(function() {
    var $sidebar = $('#sidebar');
    var $toggleSidebar = $('#toggleSidebar');

    $toggleSidebar.on('click', function(e) {
        e.preventDefault();
        $sidebar.toggle();
        $('i', $toggleSidebar).toggleClass('icon-chevron-left');
        $('i', $toggleSidebar).toggleClass('icon-chevron-right');
    });
});
