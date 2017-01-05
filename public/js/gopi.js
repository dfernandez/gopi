var connection;

function connect() {
    connection = new WebSocket('ws://localhost:8000/ws');

    // Log errors
    connection.onerror = function (error) {
        console.log(error);
    };

    // Log messages from the server
    connection.onmessage = function (e) {
        message = jQuery.parseJSON(e.data)
        switch (true) {
            case message.cmd == "button_timer_off":
                $('#toggle_lights_button').addClass('btn-danger')
                $('#toggle_lights_button').removeClass('btn-success')
                $('#toggle_lights_button').html('<i class="fa fa-play" style="font-size:16px"></i>')
                break;
            default:
                console.log(message);
        }
    };
}

$('#toggle_lights_button').click(function(){
    connection.send('button_timer_on')
    $('#toggle_lights_button').removeClass('btn-danger')
    $('#toggle_lights_button').addClass('btn-success')
    $('#toggle_lights_button').html('<i class="fa fa-refresh fa-spin" style="font-size:16px"></i>');
});