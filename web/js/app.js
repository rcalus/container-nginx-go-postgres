$(function() {
    $('#result').hide();
    $('#error').hide()
    getAverageSize();
    $('#sendSize').click(function() {
        $('#error').hide()
        $('#errorText').text('');
        $.ajax({
            contentType: 'application/json',
            data: JSON.stringify({
                "size": $("#size").val()
            }),
            dataType: 'json',
            success: function(data){
                console.log("Posted JSON size");
                getAverageSize();
                $("#size").val('');
            },
            error: function(data){
                var errorObj = $.parseJSON(data.responseText);
                $('#errorText').text(errorObj.error);
                $('#error').show();
            },
            type: 'POST',
            url: '/truetosize'
        });
    });
});


function getAverageSize(){
    $.get('/truetosize', function(data){
        $('#average').text(parseFloat(data.averageTrueToSize).toFixed(2));
        var averageText = '';
        if (data.averageTrueToSize > 1){
            averageText = 'Really small'
        }
        if (data.averageTrueToSize > 2){
            averageText = 'Small'
        }
        if (data.averageTrueToSize > 3){
            averageText = 'True to Size'
        }
        if (data.averageTrueToSize > 4){
            averageText = 'Big'
        }
        if (data.averageTrueToSize > 4){
            averageText = 'Really Big'
        }

        if (averageText !== '') {
            $('#result').show();
            $('#averageText').text(averageText);
        }
        

    });

}