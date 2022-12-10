let addUser = () => {
    let registerBtn = document.getElementById("modalRegisterUser")
    
        $.ajax({
            url: "/api/add_user",
            data: {
                "username": $("#user-username").val(),
                "password": $("#user-password").val(),
                "firstname": $("#user-fname").val(),
                "lastname": $("#user-lname").val(),
            },
            success: function() {
                $("#register-modal").removeClass("active")
            }
        })
}