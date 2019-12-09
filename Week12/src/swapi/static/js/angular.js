var app = angular.module("swapi", []);
app.controller('swapiCtrl', function($scope,$http) {
    $scope.interactive_call = function(){
        var content = jQuery('#interactive').val()
        if(content == ''){
            content = 'people/1/';
        }
        var call_url = '/api/' + content;
        jQuery.ajax({
      dataType: 'json',
      url: call_url,
      context: document.body,
      success: function(data){
            var d = jQuery.parseJSON(data);
            jQuery('#interactive_output').text(JSON.stringify(d, null, '\t'));
            
        },
       error: function(data){
       		jQuery('#interactive_output').text("404 Error");
       }
    });
   }
});


