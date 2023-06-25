PayPal SDK for Golang

## 帮助

在集成的过程中有遇到问题，欢迎加 QQ 群 203357977 讨论。

## 已实现接口

#### PaymentS API

* **Get an access token (获取 Access Token)**

	POST /v1/oauth2/token

* **Create a payment (创建账单)**

	POST /v1/payments/payment
	
* **Execute approved PayPal payment (核准账单支付信息)**

	POST /v1/payments/payment/#payment_id/execute
	
* **Show payment details (获取账单详情)**

	GET /v1/payments/payment/#payment_id 
	
* **List payments (获取账单列表)**

	GET /v1/payments/payment
	
* **Show sale details (获取交易详情)**
	
	GET /v1/payments/sale/#sale_id
	
* **Refund sale (发起退款)**

	POST /v1/payments/sale/#sale_id/refund
	
* **Show refund details (获取退款详情)**

	GET /v1/payments/refund/#refund_id
	
#### Webhooks API

* **Create webhook (创建钩子)**

	POST /v1/notifications/webhooks
	
* **Show webhook details (获取钩子详情)**

	POST /v1/notifications/webhooks/#webhook_id
	
* **List all webhooks (获取所有的钩子)**

	GET /v1/notifications/webhooks
	
* **Delete webhook (删除钩子)**

	DELETE /v1/notifications/webhooks/#webhook_id
	
## 集成流程

#### Sandbox 账户

从网页 [Sandbox - Accounts](https://developer.paypal.com/developer/accounts/) 可以创建和查看 Sandbox 环境的用户，这样就可以使用测试环境进行实际的支付和收款操作了。

大家在创建 Sandbox 账户的时候，要注意账户有两种类型，即商户和个人，实际测试的时候，应该创建一个商户账户和一个个人账户，商户账户和 App 进行关联，这样个人账户向 App 支付的时候，款项将转入商户的账户中。

Sandbox 账户有专门测试网站 [https://www.sandbox.paypal.com/cn/](https://www.sandbox.paypal.com/cn/)，该网站提供的功能和网站 [https://www.paypal.com/cn/](https://www.paypal.com/cn/) 提供的功能是一致的，比如收入、支出、退款操作等，只不过里面的数据都是测试数据。

#### 创建 App

访问网站 [https://developer.paypal.com/](https://developer.paypal.com/), 使用 PayPal 的账号登录，进入 [My Apps & Credentials](https://developer.paypal.com/developer/applications/) 页面, 找到 **REST API apps**，点击 **Create App** 创建一个新的 App，创建成功之后可以获取到 **Client ID** 和 **Secret**，我们后续在进行认证的时候会用到这两个参数。

创建 App 的时候，需要关联一个 Sanbox 环境的商户账户。

#### 获取 Access Token 

```Golang

import "github.com/smartwalle/paypal"

var client = paypal.New("ClientID", "Secret", false) // 第三个参数用于标记是否为生产环境，true 为生产环境，false 为 Sandbox 环境。

var token, err = client.GetAccessToken() // 获取 Access Token

```

在实际使用过程中，一般不需要单独调用此方法获取 Access Token，除非你有需要。

在访问其它需要认证的接口的时候，组件会自动判断当前是否有正常可用的 Access Token，如果没有，会先向 PayPal 请求 Access Token, 然后再进行业务接口的访问。

#### 创建账单

创建账单提供了两种方式：

##### 1. 快速创建账单

```Golang
var payment, err = client.ExpressCreatePayment(invoiceNumber, total, currency, cancelURL, returnURL)
...
```

##### 2. 高级接口

```Golang
var p = &paypal.Payment{}
p.Intent = paypal.PaymentIntentSale
p.Payer = &paypal.Payer{}
p.Payer.PaymentMethod = "paypal"
p.RedirectURLs = &paypal.RedirectURLs{}
p.RedirectURLs.CancelURL = "http://www.baidu.com"
p.RedirectURLs.ReturnURL = "http://127.0.0.1:9001/paypal"

var transaction = &paypal.Transaction{}
p.Transactions = []*paypal.Transaction{transaction}

transaction.Amount = &paypal.Amount{}
transaction.Amount.Total = "30.11"
transaction.Amount.Currency = "USD"
transaction.Amount.Details = &paypal.AmountDetails{}
transaction.Amount.Details.Subtotal = "30.00"
transaction.Amount.Details.Tax = "0.07"
transaction.Amount.Details.Shipping = "0.03"
transaction.Amount.Details.HandlingFee = "1.00"
transaction.Amount.Details.ShippingDiscount = "-1.00"
transaction.Amount.Details.Insurance = "0.01"

transaction.Description = "This is the payment transaction description."
transaction.Custom = "EBAY_EMS_90048630024435"
transaction.InvoiceNumber = uuid.New() // 随机生成一串 Invoice Number

transaction.PaymentOptions = &paypal.PaymentOptions{}
transaction.PaymentOptions.AllowedPaymentMethod = "INSTANT_FUNDING_SOURCE"
transaction.SoftDescriptor = "ECHI5786786"

transaction.ItemList = &paypal.ItemList{}
transaction.ItemList.ShippingAddress = &paypal.ShippingAddress{}
transaction.ItemList.ShippingAddress.RecipientName = "Hello World"
transaction.ItemList.ShippingAddress.Line1 = "4thFloor"
transaction.ItemList.ShippingAddress.Line2 = "unit#34"
transaction.ItemList.ShippingAddress.City = "SAn Jose"
transaction.ItemList.ShippingAddress.CountryCode = "US"
transaction.ItemList.ShippingAddress.PostalCode = "95131"
transaction.ItemList.ShippingAddress.Phone = "011862212345678"
transaction.ItemList.ShippingAddress.State = "CA"

var i1, i2 = &paypal.Item{}, &paypal.Item{}
transaction.ItemList.Items = []*paypal.Item{i1, i2}

i1.Name = "hat"
i1.Description = "Brown color hat"
i1.Quantity = "5"
i1.Price = "3"
i1.Tax = "0.01"
i1.SKU = "1"
i1.Currency = "USD"

i2.Name = "handbag"
i2.Description = "Black color hand bag"
i2.Quantity = "1"
i2.Price = "15"
i2.Tax = "0.02"
i2.SKU = "product34"
i2.Currency = "USD"

p.NoteToPayer = "Contact us for any questions on your order."

var payment, err = client.CreatePayment(p)
fmt.Println(payment, err)

```

这种创建账单的方式虽然会比较复杂，但是信息也比较全面，用户在支付的时候，也能够看到详细的商品清单信息，会显得更加的友好。

#### 支付

从创建账单返回的 Payment 信息中，我们会得到一个 Link 类型的数组 payment.Links，该数组一般含有三个数据，我们要取出其中一个 **rel** 属性的值为 **approval_url** 的 Link，然后在浏览器里面打开其 href 属性对应的链接地址，这样就可以进行登录并支付了。

OK，PayPal 显示支付成功了，也成功跳转到了我们设置的 ReturnURL，这时候进入商户账户，但是还是没有任何的入账信息，这是为什么呢？

创建账单的时候，Payment 有一个 RedirectURLs 属性，该属性有两个值，一个为 CancelURL，另一个为 ReturnURL。

CancelURL 为用户取消支付跳转的 URL；ReturnURL 为用户支付成功跳转的 URL；

当我们支付成功之后，浏览器将跳转到 ReturnURL，跳转到该 URL 之后，我们还要进行一步工作，这样才能保证支付完成。

#### 核准账单支付信息

支付成功之后，浏览器跳转到 ReturnURL 的时候，附带了几个参数：paymentId、token 和 PayerID。

```
http://192.168.192.250:3000/paypal?paymentId=PAY-37A82711YL064934DLB4G3AQ&token=EC-0DG12278CE6129333&PayerID=XV9HF9K25FB38
```

核准账单支付信息的时候，需要用到 paymentId 和 payerID。

我们在提供的 ReturnURL 接口中应执行核准账单支付信息的操作。

```Golang
var payment, err = client.ExecuteApprovedPayment(paymentId, payerID)
...

```

如果返回的 payment 的 State 为 “approved”，则表示核准账单支付成功（注意：不能以此来判断是否支付成功，即实际到账）。

**如果要判断是否实际到账，需要判断返回结果中的 transactions[xx].related_resources[xx].sale.state 的值，当该字段的值为 completed 的时候，才能说明已到账。**

到此，可以算是完成了一个完整的收款流程，如果想要更加严谨，还需要加上 webhook。

#### 总结

一个简单的支付流程：

1. 初始化 paypal 信息:
	
	```
	var client = paypal.New(...)
	```
	
2. 创建账单

	```
	var p = &Payment{}
	...
	client.CreatePayment(p)
	```

3. 从账单中获取类型为 approval_url 的 URL，浏览器打开进行支付
4. ReturnURL 接口中进行核准账单支付

	```
	var client = paypal.New(...)
	var p, err = client.ExecuteApprovedPayment(paymentId, payerID)
	...
	```
	

## Webhook (钩子)

待续...