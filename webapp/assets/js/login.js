$('#login').on('submit', fazerLogin);

function fazerLogin(evento) {
    evento.preventDefault();
    //console.log("Dentro da funcao fazer login");

   $.ajax({
        url: "/login",
        method: "POST",
        data: {
            email: $('#email').val(),
            senha: $('#senha').val(),
        }
    }).done(function() {
        window.location = "/home";
        alert("chegou aqui!");
    }).fail(function() {
        alert("Usuário ou senha invalidos!");
    });
}