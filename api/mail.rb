require 'sendgrid-ruby'
include SendGrid

Handler = Proc.new do |req, res|
  body = eval(req.body)
  name = body[:name]
  last_name = body[:last_name]
  email = body[:email]
  subject = body[:subject]
  message = body[:message]

  # Send user email to me
  from = Email.new(email: 'ifrancisco.iglesias@gmail.com')
  to = Email.new(email: email)
  content = Content.new(type: 'text/html', value: message)
  mail = Mail.new(from, subject, to, content)
  sg = SendGrid::API.new(api_key: ENV['SENDGRID_API_KEY'])
  sg.client.mail._('send').post(request_body: mail.to_json)

  # Send auto-response
  automatic_response = Mail.new
  automatic_response.from = to
  personalization = Personalization.new
  personalization.add_to(from)
  personalization.add_dynamic_template_data({
    "name" => name,
    "last_name" => last_name
  })
  automatic_response.add_personalization(personalization)
  automatic_response.template_id = 'd-875a38a32cd443188b520a0cf70fc7f0'  
  sg.client.mail._('send').post(request_body: automatic_response.to_json)
  res.status = 200
  res.body = ''
end
