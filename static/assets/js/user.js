let addUser = () => {
        $.ajax({
            url: "/api/add_user",
            data: {
                "username": $("#user-username").val(),
                "password": $("#user-password").val(),
                "firstname": $("#user-fname").val(),
                "lastname": $("#user-lname").val(),
            },
            success: function() {
                $("#register-modal").empty()
                $("#register-modal").append(`
                    <h2>Registered Successfully</h2>
                    <button onclick="closeRegModal()">Close</button>    
                `)
              
            }
        })
}

let closeRegModal = () => {
    $("#register-modal").removeClass("active")
}
