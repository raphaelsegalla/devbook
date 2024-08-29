$('#nova-publicacao').on('submit', criarPublicacao);
$('.curtir-publicacao').on('click', curtirPublicacao);

function criarPublicacao(evento) {
    evento.preventDefault();

    $.ajax({
        url: "/publicacoes",
        method: "POST",
        data: {
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val()
        }
    }).done(function() {
        window.location = "/home";
    }).fail(function() {
        alert("Erro ao criar publicacao!");
    }) 
}

function curtirPublicacao(evento) {
    evento.preventDefault();

    const elementoClicado = $(evento.target);
    const publicacaoId = elementoClicado.closest('div').data('publicacao-id');

    $.ajax({
        url: `/publicacoes/${publicacaoId}/curtir`,
        method: "POST"
    }).done(function(evento) {
        console.log(evento)
        alert("Publicaçâo curtida!");
    }).fail(function(evento) {
        console.log(evento)
        alert("Erro ao curtir publicaçâo!");
    })
}