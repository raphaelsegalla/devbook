$('#nova-publicacao').on('submit', criarPublicacao);
// $('.curtir-publicacao').on('click', curtirPublicacao);
$(document).on('click', '.curtir-publicacao', curtirPublicacao);
$(document).on('click', '.descurtir-publicacao', descurtirPublicacao);

$('#atualizar-publicacao').on('click', atualizarPublicacao);
$('.deletar-publicacao').on('click', deletarPublicacao);

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

    elementoClicado.prop('disabled', true)
    $.ajax({
        url: `/publicacoes/${publicacaoId}/curtir`,
        method: "POST"
    }).done(function(evento) {
        const contadorDeCurtidas = elementoClicado.next('span');
        const quantidadeDeCurtidas = parseInt(contadorDeCurtidas.text());

        contadorDeCurtidas.text(quantidadeDeCurtidas + 1);

        elementoClicado.addClass('descurtir-publicacao')
        elementoClicado.addClass('text-danger')
        elementoClicado.removeClass('curtir-publicacao')
    }).fail(function(evento) {
        alert("Erro ao curtir publicaçâo!");
    }).always(function() {
        elementoClicado.prop('disabled', false)
    })
}

function descurtirPublicacao(evento) {
    evento.preventDefault();

    const elementoClicado = $(evento.target);
    const publicacaoId = elementoClicado.closest('div').data('publicacao-id');

    elementoClicado.prop('disabled', true)
    $.ajax({
        url: `/publicacoes/${publicacaoId}/descurtir`,
        method: "POST"
    }).done(function(evento) {
        const contadorDeCurtidas = elementoClicado.next('span');
        const quantidadeDeCurtidas = parseInt(contadorDeCurtidas.text());

        contadorDeCurtidas.text(quantidadeDeCurtidas - 1);

        elementoClicado.removeClass('descurtir-publicacao')
        elementoClicado.removeClass('text-danger')
        elementoClicado.addClass('curtir-publicacao')
    }).fail(function(evento) {
        alert("Erro ao descurtir publicaçâo!");
    }).always(function() {
        elementoClicado.prop('disabled', false)
    })
}

function atualizarPublicacao(evento) {
    $(this).prop('disabled', true);

    const publicacaoId = $(this).data('publicacao-id');

    $.ajax({
        url: `/publicacoes/${publicacaoId}`,
        method: "PUT",
        data: {
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val()
        }
    }).done(function(evento) {
        alert("Publicação editada com sucesso!");
    }).fail(function(evento) {
        alert("Erro ao editar publicaçâo!");
    }).always(function() {
        $('#atualizar-publicacao').prop('disabled', false)
    })
}

function deletarPublicacao(evento) {
    evento.preventDefault();

    const elementoClicado = $(evento.target);
    const publicacao = elementoClicado.closest('div')
    const publicacaoId = publicacao.data('publicacao-id');

    elementoClicado.prop('disabled', true)

    $.ajax({
        url: `/publicacoes/${publicacaoId}`,
        method: "DELETE"
    }).done(function() {
        publicacao.fadeOut("slow", function() {
            $(this).remove();
        });
    }).fail(function() {
        alert("Erro ao excluir a publicação!")
    });
}