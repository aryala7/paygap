package zarinpal

// providerCodeMessages zarinpal code message list
var providerCodeMessages = map[int]string{
	100: "عملیات موفق",
	101: "تراکنش وریفای شده",
	-9:  "خطای اعتبار سنجی",
	-10: "ای پی و يا مرچنت كد پذيرنده صحيح نيست",
	-11: "مرچنت کد فعال نیست لطفا با تیم پشتیبانی ما تماس بگیرید",
	-12: "تلاش بیش از حد در یک بازه زمانی کوتاه",
	-15: "ترمینال شما به حالت تعلیق در آمده با تیم پشتیبانی تماس بگیرید",
	-16: "سطح تاييد پذيرنده پايين تر از سطح نقره اي است",
	-30: "اجازه دسترسی به تسویه اشتراکی شناور ندارید",
	-32: "Wages is not valid, Total wages(floating) has been overload max amount",
	-33: "درصد های وارد شده درست نیست",
	-34: "مبلغ از کل تراکنش بیشتر است",
	-35: "تعداد افراد دریافت کننده تسهیم بیش از حد مجاز است",
	-40: "Invalid extra params, expire_in is not valid",
	-50: "مبلغ پرداخت شده با مقدار مبلغ در وریفای متفاوت است",
	-52: "خطای غیر منتظره با پشتیبانی تماس بگیرید",
	-53: "اتوریتی برای این مرچنت کد نیست",
	-54: "اتوریتی نامعتبر است",
}

type Zarinpal struct {
	merchantID string
	sandbox    bool
}

type request struct {
	MerchantID  string `json:"merchant_id"`
	Amount      int    `json:"amount"`
	CallBackURL string `json:"call_back_url"`
	Description string `json:"description"`
	Email       string `json:"email"`
	Mobile      string `json:"mobile"`
}

type response struct {
	Status    int    `json:"status"`
	Authority string `json:"authority"`
}
