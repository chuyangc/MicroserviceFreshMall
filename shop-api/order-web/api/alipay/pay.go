package main

import (
	"fmt"
	"github.com/smartwalle/alipay/v3"
)

func main() {
	appID := "2021000122611696"
	privateKey := "MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQDwSQuA8UFCRwE9pK6wjpzanl+AejnkwdSkbv5cBljgwZ3iUsh6y/6y6EcYQ4+piLBOOCrg5mSaqLYBpz7/iLs9msNTzWq5OE7nwrk1bEbo0Rj+1WYPUFZQhQL6CkfCCb4kVaHxeZzhE0cHo9xSJWRVi+caEF9IsM6VJBJu2YXvT2tCPwaLCCa5nRJzzkbockZ017uwIKBHbhLpnX9GmKB20yzoSViPSBaeEGdCWG4mKRncSVUhq8dit7pvBdrst2wWbjjMk/RPD4g5MbSIFkuc7MMDjesRP90yjBwUQtDCPHXs+hWukJ/g56+ax3CJZHRxFtdBNqBZ7yKQztte8sEVAgMBAAECggEARQB08T5WwzWowY79K26I1K8ONdLjtTGEYwQMv1iDRWfUcx3avIjAR5g0cl9Ubhb2qj+u8I647UDto2Pnz3HwcyxyUyp2L2JgJmXg0dqaMll5mBSoDlW/s7e+txckrDAoDj8ZFkMaLfhfOW5w4pYiTf6zCuUQt8suR93n/TUyJRpW++tYZVZKgsXvqIUTwxj2JeLEYvH9puxtgrLX58t5QHoB0f+rxSVD88W4tBY4Ml2KJr/NbValrWsFGpvt8jMUZU8TQVzmreR05BG7jFnRNWoSlLlsIKNg4wAwn7baxPHH5WLsI8GcRDOXZQh6nTkohTtkVR1fodRDNAGHDO6HsQKBgQD5U+S1zAHxe/vGEx1DkyCHJeR9WixEemMLxLOzBWY2qBTo5+651HPUHgVfSi7l5dI0tjz2tcT236PtMqfQLwjM0VlR/jXEH7m0BHBnA0kWpb56GP/5YB1V2GoJ6vKf+3mGDVlLcERnc1hcivAu3kpVGauREJN3TtWmqL1NnJkAWwKBgQD2tzQde4TggnPxC0HnT0UbcMfm7f+q9JItTDFP1koRSGodwD0NHfNQJJx2j4vLDyEFoKmOuOQbfQdAx+pbupEYtUhdr7ucJrIgC/dWW/mxd29343spN3r6nWyCW0TW5LvoGq3h0ikuELzs+tcZHA+LConixpGTXyoshLSiSvP/TwKBgE+jGUlsKS694FSLJGzCIMCqPMpBNCSHRv2qTY+f6N8KXutpsZnPn1OgZyzhoAs0lijaEKzosEn+cvi/llRrwY7SS4ph/UBwtbsnM9Pje8PtGuMa+x/nMFeMMYqLbgXlqBJGT0BGUsMMV6vvgPonbGy0L1W9iqywFJQQD68rlr9DAoGAR3Fx8/+q0OC19l1OLk90MagNG0BcQwMjuV5RAU+Kj0qrAAaFJ2E+7jxL2sFit+CfrWOC9kNwOs2P5iB+KyXxkngchpS2/VbvSfxtGWL1AYEWlF8ZcSwRvrULkQwg+SGvkFz5cWVMa3yJWQ6ibzEDBz58A2GBEm4CZnXfYQfXdF0CgYBORPbBjEptKSgApaC+/2VfR+/m0ED9AxZ/f6B02TZ6oNUQ10EIuR+YVFdo8kPxKr4+q7lAAUIJoH1zDncojHQmJ5OrI0Git3NsJt0Bf1xBtm/sgC1eJFNd6OPsEqhi1qjmY8T9uCKNfXito+1as1zrkFxzIjpoyV1c6bwRUHBNEQ=="
	aliPublicKey := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArCds25gWOM4lqL77068qIAUmoRbfdkrR2Sz7xuFpHf7XHcmnmxH1cYYCK68A+OD7ni/jOf6Qh4y/g3g62aMNF6rLQYointroCh6hvKQGJFHikEw3N71I0FAhakg3RFgJ1wekw2ZhkcZPUvL7fK+juGrD++pJskxh60SRt3NMsr07n34foSG0eVTUXS5unGS9JWtUlwGsyYvJmIn7AOB6WYPvJJB0/ju3wfwOmOD7VjQfO+4tMFXzKrkPS15v2KYw4myakKrXuGMfbVpzzbIixvrmOvtpzlZvR2LElQ27MKGPgf5p+/IP9i1RKTfPQtbGqCmoBxk72YDc7gWGvrst/wIDAQAB"
	var client, err = alipay.New(appID, privateKey, false)
	if err != nil {
		panic(err)
	}

	err = client.LoadAliPayPublicKey(aliPublicKey)
	if err != nil {
		panic(err)
	}

	var p = alipay.TradePagePay{}
	p.NotifyURL = "http://ecf6617734fb7c30.natapp.cc/o/v1/pay/alipay/notify"
	p.ReturnURL = "http://127.0.0.1:8089"
	p.Subject = "生鲜订单支付"
	p.OutTradeNo = "chuyangc"
	p.TotalAmount = "10.00"
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"

	url, err := client.TradePagePay(p)
	if err != nil {
		panic(err)
	}
	fmt.Println(url.String())
}
