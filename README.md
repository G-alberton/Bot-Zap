# Bot-Zap
Feito para integração com a meta API oficial Whatsapp
A linguagem escolhida foi GO, visto sua facil utilizadade para integração com APIs, visto o pouco tempo que tinhamos para criar a aplicação.

Segue as templates que são utilizadas:

 -> Primeiro_contato_v3:

        Olá, {{Nome}}.

        Em nome da (Empresa), informamos que a parcela no valor de {{Valor}} possui vencimento em {{DataVencimento}}.

        Caso o boleto não tenha sido recebido, permanecemos à disposição para reenvio.

 -> Avisa_vencimento_v3:
        Olá, {{Nome}}.

        Consta em nosso sistema a pendência referente à parcela do contrato firmado com a (EMPRESA), no valor de {{Valor}}, com vencimento em {{DataVencimento}}.

        Informamos que, após o prazo de 5 (cinco) dias do vencimento, foram encaminhadas instruções para protesto junto ao CPF vinculado ao contrato.

        Para regularização da situação, o pagamento deverá ser realizado diretamente no cartório responsável ou junto à Mascor.

        Em caso de dúvidas, permanecemos à disposição para esclarecimentos.

 -> Manda_boleto_v3:
        cabeçalho: "TIPO PDF"

        Olá, {{Nome}}.

        Encaminhamos, em anexo a esta mensagem, a imagem referente ao boleto da parcela vinculada ao seu contrato com a (Empresa), no valor de {{Valor}}, com vencimento em {{DataVencimento}}.

        Solicitamos que verifique as informações constantes no documento. Em caso de dúvidas ou necessidade de esclarecimentos, permanecemos à disposição.         

 -> Envia_Pix_v2:
        cabeçalho: "TIPO IMAGEM"

        Olá, {{Nome}}

        Segue a imagem do QR Code para que você possa realizar o pagamento.

        Caso prefira, utilize a chave Copia e Cola abaixo:
        {{ChavePIX}}

        Tenha um bom dia.


!!Lembrando, as variaveis, devem ser numero e não nome, pode ocorrer erro no processo!!