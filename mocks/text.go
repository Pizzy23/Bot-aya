package mocks

const (
	InvestIntro = "Legal! Aqui estão os tipos de investimento disponíveis:\n" +
		"1 - Renda Fixa\n" +
		"2 - Renda Variável\n" +
		"3 - Tesouro Direto\n" +
		"4 - CDB PagBank\n" +
		"5 - Fundos de Investimento\n\n" +
		"Digite o número da opção que deseja explorar."

	InvestInvalidOption = "Número inválido. Por favor, tente novamente digitando o número do tipo de investimento."

	InvestNotFound      = "Nenhum investimento encontrado para o tipo: %s."
	InvestSummaryFormat = "Resumo do investimento em %s:\n\nTotal Investido: R$%.2f\n\nRendimentos: +R$%.2f\n"
	InvestDetailsPrompt = "Deseja mais informações sobre algum investimento específico?\nDigite o número da opção ou escreva 'não' para encerrar."

	InvestExitPrompt = "Fico feliz em ajudar! Caso precise de atualizações ou qualquer outra assistência, é só chamar. Boa semana! 📈"
	InvestExitOption = "Número inválido. Por favor, tente novamente digitando o número do investimento específico. \n\nDigite 'não' para sair."

	SlipsIntro          = "Olá! Aqui estão os boletos em aberto:\n"
	SlipsNoPending      = "Olá! Não há boletos pendentes no momento."
	SlipsPrompt         = "\nPor favor, digite o número do boleto que deseja agendar para pagamento."
	InvalidSlipNumber   = "Número inválido. Por favor, tente novamente digitando o número do boleto."
	SlipNotFound        = "Boleto não encontrado. Por favor, tente novamente."
	SlipDetailsTemplate = "Obrigado! Aqui estão os detalhes do boleto:\n\nNome: %s\nValor: R$%.2f\nCódigo de Barras: %s\n\nConfirma o agendamento?"
	PaymentConfirmed    = "Perfeito! Seu pagamento foi agendado. Vou enviar uma notificação de confirmação no dia do pagamento. Algo mais em que posso ajudar?"
	PaymentCancelled    = "Cancelando agendamento. Posso ajudar com mais alguma coisa?"

	RechargePrompt       = "Por favor, informe o número do celular (11 dígitos) ou do Bilhete Único (10 dígitos):"
	InvalidNumber        = "Número inválido. Por favor, digite um número de celular (11 dígitos) ou Bilhete Único (10 dígitos)."
	RechargeValuePrompt  = "Obrigado! Para qual valor você deseja recarregar o %s?"
	InvalidRechargeValue = "Valor inválido. Por favor, digite um valor numérico positivo."
	RechargeConfirmation = "Confirmando, uma recarga de R$%.2f para o %s %s. Está correto?"
	RechargeSuccess      = "Prontinho! Sua recarga de R$%.2f foi realizada com sucesso para o %s %s. Posso ajudar em mais alguma coisa?"
	RechargeCancellation = "Cancelando a recarga. Posso ajudar com mais alguma coisa?"

	UnrecognizedCommand     = "Agora nos informe o menu que deseja acessar:\n\n- Investimento\n- Agenda Integrada\n- Recarga"
	EmailPrompt             = "Ola! Sou a Aia, sua assistente financeira.\nInforme seu e-mail para iniciarmos a jornada."
	UserCreationSuccess     = "Perfeito! \nAgora nos informe o menu que deseja acessar:\n\n- Investimento\n- Agenda Integrada\n- Recargas"
	WeeklySendCodeReceived  = "Código de comando recebido. Forçando o envio da mensagem semanal."
	NoPendingSlips          = "Olá! Não há boletos pendentes no momento."
	ChoosePaymentOption     = "Por favor, digite o número do boleto que deseja agendar para pagamento."
	InvalidPaymentOption    = "Número inválido. Por favor, tente novamente digitando o número do boleto."
	PaymentScheduledSuccess = "Perfeito! Seu pagamento foi agendado. Vou enviar uma notificação de confirmação no dia do pagamento. Algo mais em que posso ajudar?"
	CancelPaymentSchedule   = "Cancelando agendamento. Posso ajudar com mais alguma coisa?"
	WelcomeMessage          = "Olá! Sou a AIA, sua assistente financeira.\nAgora nos informe o menu que deseja acessar:\n\n- Investimento\n- Agenda Integrada\n- Recargas"
)
