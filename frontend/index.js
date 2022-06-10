$(".creds-btn").button({
  inline: "true",
});
function getResponse() {
  $.getJSON("https://t32m4gek4d.execute-api.us-east-1.amazonaws.com/dev/get_response", function (data) {
    $('#responseField').html(data.message);
    console.log(data.message);
  })
}

function send_form() {
  $.post("http://34.238.126.34:80/post_data", { "first": $('#first').val(), "last": $('#last').val() },
    function (data) { console.log(data.message) },
    "json").fail(function (response) { console.log("error"); console.log(response) })
  return false
}