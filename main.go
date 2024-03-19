package main

import (
	"log"
    "net/http"
    "github.com/gin-gonic/gin"
)


/**
* Main function. Handles all of the routes...
 */
func main() {

    log.Println("Server is starting ....... ")

    ginMode := "release"
    gin.SetMode(ginMode)
    router := gin.New()
    router.Use(gin.Recovery())

    router.GET("/", HomepageHandler)
 	 //
    router.GET("/api/100", ServeHTTP100)
    router.GET("/api/101", ServeHTTP101)
    router.GET("/api/102", ServeHTTP102)
    router.GET("/api/103", ServeHTTP103)
    //
    router.GET("/api/200", ServeHTTP200)
    router.GET("/api/201", ServeHTTP201)
    router.GET("/api/202", ServeHTTP202)
    router.GET("/api/203", ServeHTTP203)
    router.GET("/api/204", ServeHTTP204)
    router.GET("/api/205", ServeHTTP205)
    router.GET("/api/206", ServeHTTP206)
    router.GET("/api/207", ServeHTTP207)
    router.GET("/api/208", ServeHTTP208)
    router.GET("/api/226", ServeHTTP226)
    //
    router.GET("/api/300", ServeHTTP300)
    router.GET("/api/301", ServeHTTP301)
    router.GET("/api/302", ServeHTTP302)
    router.GET("/api/303", ServeHTTP303)
    router.GET("/api/304", ServeHTTP304)
    router.GET("/api/305", ServeHTTP305)
    //
    router.GET("/api/307", ServeHTTP307)
    router.GET("/api/308", ServeHTTP308)
    //
    router.GET("/api/400", ServeHTTP400)
    router.GET("/api/401", ServeHTTP401)
    router.GET("/api/402", ServeHTTP402)
    router.GET("/api/403", ServeHTTP403)
    router.GET("/api/404", ServeHTTP404)
    router.GET("/api/405", ServeHTTP405)
    router.GET("/api/406", ServeHTTP406)
    router.GET("/api/407", ServeHTTP407)
    router.GET("/api/408", ServeHTTP408)
    router.GET("/api/409", ServeHTTP409)
    router.GET("/api/410", ServeHTTP410)
    router.GET("/api/411", ServeHTTP411)
    router.GET("/api/412", ServeHTTP412)
    router.GET("/api/413", ServeHTTP413)
    router.GET("/api/414", ServeHTTP414)
    router.GET("/api/415", ServeHTTP415)
    router.GET("/api/416", ServeHTTP416)
    router.GET("/api/417", ServeHTTP417)
    router.GET("/api/418", ServeHTTP418)
    //
    router.GET("/api/421", ServeHTTP421)
    router.GET("/api/422", ServeHTTP422)
    router.GET("/api/423", ServeHTTP423)
    router.GET("/api/424", ServeHTTP424)
    router.GET("/api/425", ServeHTTP425)
    router.GET("/api/426", ServeHTTP426)
    router.GET("/api/428", ServeHTTP428)
    router.GET("/api/429", ServeHTTP429)
    //
    router.GET("/api/431", ServeHTTP431)
    router.GET("/api/451", ServeHTTP451)
    //
    router.GET("/api/500", ServeHTTP500)
    router.GET("/api/501", ServeHTTP501)
    router.GET("/api/502", ServeHTTP502)
    router.GET("/api/503", ServeHTTP503)
    router.GET("/api/504", ServeHTTP504)
    router.GET("/api/505", ServeHTTP505)
    router.GET("/api/506", ServeHTTP506)
    router.GET("/api/507", ServeHTTP507)
    router.GET("/api/508", ServeHTTP508)
    //
    router.GET("/api/510", ServeHTTP510)
    router.GET("/api/511", ServeHTTP511)

    router.Run()
}


func HomepageHandler(c *gin.Context) {
	    c.JSON(http.StatusOK, gin.H{"message":"Welcome to the Tech Company listing API with Golang"})
}

//100 StatusContinue
 func ServeHTTP100(c *gin.Context) {
    c.JSON(http.StatusContinue, nil)
 }


//101 StatusSwitchingProtocols
 func ServeHTTP101(c *gin.Context) {
    c.JSON(http.StatusSwitchingProtocols, nil)
 }

//102 StatusProcessing
 func ServeHTTP102(c *gin.Context) {
    c.JSON(http.StatusProcessing, nil)
 }

// //103 StatusEarlyHints
 func ServeHTTP103(c *gin.Context) {
    c.JSON(http.StatusEarlyHints, nil)
 }

//200 StatusOK
 func ServeHTTP200(c *gin.Context) {
    c.JSON(http.StatusOK, nil)
 }



//201 StatusCreated
 func ServeHTTP201(c *gin.Context) {
    c.JSON(http.StatusCreated, nil)
 }


//202 StatusAccepted
 func ServeHTTP202(c *gin.Context) {
    c.JSON(http.StatusAccepted, nil)
 }


//203 StatusNonAuthoritativeInfo
 func ServeHTTP203(c *gin.Context) {
    c.JSON(http.StatusNonAuthoritativeInfo, nil)
 }


//204 StatusNoContent
 func ServeHTTP204(c *gin.Context) {
    c.JSON(http.StatusNoContent, nil)
 }


//205 StatusResetContent
 func ServeHTTP205(c *gin.Context) {
    c.JSON(http.StatusResetContent, nil)
 }


//206 StatusPartialContent
 func ServeHTTP206(c *gin.Context) {
    c.JSON(http.StatusPartialContent, nil)
 }


//207 StatusMultiStatus
 func ServeHTTP207(c *gin.Context) {
    c.JSON(http.StatusMultiStatus, nil)
 }


//208 StatusAlreadyReported
 func ServeHTTP208(c *gin.Context) {
    c.JSON(http.StatusAlreadyReported, nil)
 }

//226 StatusIMUsed
 func ServeHTTP226(c *gin.Context) {
    c.JSON(http.StatusIMUsed, nil)
 }


//300 StatusMultipleChoices
 func ServeHTTP300(c *gin.Context) {
    c.JSON(http.StatusMultipleChoices, nil)
 }

//301 StatusMovedPermanently
 func ServeHTTP301(c *gin.Context) {
    c.JSON(http.StatusMovedPermanently, nil)
 }

 //302 StatusFound
 func ServeHTTP302(c *gin.Context) {
    c.JSON(http.StatusFound, nil)
 }

 //303 StatusSeeOther
 func ServeHTTP303(c *gin.Context) {
    c.JSON(http.StatusSeeOther, nil)
 }

 //304 StatusNotModified
 func ServeHTTP304(c *gin.Context) {
    c.JSON(http.StatusNotModified, nil)
 }

 //305 StatusUseProxy
 func ServeHTTP305(c *gin.Context) {
    c.JSON(http.StatusUseProxy, nil)
 }

 //306 NOT USED 
 //func ServeHTTP306(c *gin.Context) {
 //   c.JSON(http.Status_NOT_Used, nil)
 //}

 //307 StatusTemporaryRedirect
 func ServeHTTP307(c *gin.Context) {
    c.JSON(http.StatusTemporaryRedirect, nil)
 }

 //308 StatusPermanentRedirect
 func ServeHTTP308(c *gin.Context) {
    c.JSON(http.StatusPermanentRedirect, nil)
 }

 //400 StatusBadRequest
 func ServeHTTP400(c *gin.Context) {
    c.JSON(http.StatusBadRequest, nil)
 }

 //401 StatusUnauthorized
 func ServeHTTP401(c *gin.Context) {
    c.JSON(http.StatusUnauthorized, nil)
 }

 //402 StatusPaymentRequired
 func ServeHTTP402(c *gin.Context) {
    c.JSON(http.StatusPaymentRequired, nil)
 }

 //403 StatusForbidden
 func ServeHTTP403(c *gin.Context) {
    c.JSON(http.StatusForbidden, nil)
 }

 //404 StatusNotFound
 func ServeHTTP404(c *gin.Context) {
    c.JSON(http.StatusNotFound, nil)
 }


 //405 StatusMethodNotAllowed
 func ServeHTTP405(c *gin.Context) {
    c.JSON(http.StatusMethodNotAllowed, nil)
 }


 //406 StatusNotAcceptable
 func ServeHTTP406(c *gin.Context) {
    c.JSON(http.StatusNotAcceptable, nil)
 }


 //407 StatusProxyAuthRequired
 func ServeHTTP407(c *gin.Context) {
    c.JSON(http.StatusProxyAuthRequired, nil)
 }

 //408 StatusRequestTimeout
 func ServeHTTP408(c *gin.Context) {
    c.JSON(http.StatusRequestTimeout, nil)
 }


 //409 StatusConflict
 func ServeHTTP409(c *gin.Context) {
    c.JSON(http.StatusConflict, nil)
 }


 //410 StatusGone
 func ServeHTTP410(c *gin.Context) {
    c.JSON(http.StatusGone, nil)
 }


 //411 StatusLengthRequired
 func ServeHTTP411(c *gin.Context) {
    c.JSON(http.StatusLengthRequired, nil)
 }

 //412 StatusPreconditionFailed
 func ServeHTTP412(c *gin.Context) {
    c.JSON(http.StatusPreconditionFailed, nil)
 }

 //413 StatusRequestEntityTooLarge
 func ServeHTTP413(c *gin.Context) {
    c.JSON(http.StatusRequestEntityTooLarge, nil)
 }

 //414 StatusRequestURITooLong
 func ServeHTTP414(c *gin.Context) {
    c.JSON(http.StatusRequestURITooLong, nil)
 }

 //415 StatusUnsupportedMediaType
 func ServeHTTP415(c *gin.Context) {
    c.JSON(http.StatusUnsupportedMediaType, nil)
 }

 //416 StatusRequestedRangeNotSatisfiable
 func ServeHTTP416(c *gin.Context) {
    c.JSON(http.StatusRequestedRangeNotSatisfiable, nil)
 }

 //417 StatusExpectationFailed
 func ServeHTTP417(c *gin.Context) {
    c.JSON(http.StatusExpectationFailed, nil)
 }

 //418 StatusTeapot
 func ServeHTTP418(c *gin.Context) {
    c.JSON(http.StatusTeapot, nil)
 }



 //421 StatusMisdirectedRequest
 func ServeHTTP421(c *gin.Context) {
    c.JSON(http.StatusMisdirectedRequest, nil)
 }


 //422 StatusUnprocessableEntity
 func ServeHTTP422(c *gin.Context) {
    c.JSON(http.StatusUnprocessableEntity, nil)
 }

 //423 StatusLocked
 func ServeHTTP423(c *gin.Context) {
    c.JSON(http.StatusLocked, nil)
 }


 //424 StatusFailedDependency
 func ServeHTTP424(c *gin.Context) {
    c.JSON(http.StatusFailedDependency, nil)
 }


 //425 StatusTooEarly
 func ServeHTTP425(c *gin.Context) {
    c.JSON(http.StatusTooEarly, nil)
 }

 //426 StatusUpgradeRequired
 func ServeHTTP426(c *gin.Context) {
    c.JSON(http.StatusUpgradeRequired, nil)
 }

 //428 StatusPreconditionRequired
 func ServeHTTP428(c *gin.Context) {
    c.JSON(http.StatusPreconditionRequired, nil)
 }


 //429 StatusTooManyRequests
 func ServeHTTP429(c *gin.Context) {
    c.JSON(http.StatusTooManyRequests, nil)
 }


 //431 StatusRequestHeaderFieldsTooLarge
 func ServeHTTP431(c *gin.Context) {
    c.JSON(http.StatusRequestHeaderFieldsTooLarge, nil)
 }



 //451 StatusUnavailableForLegalReasons
 func ServeHTTP451(c *gin.Context) {
    c.JSON(http.StatusUnavailableForLegalReasons, nil)
 }



 //500 StatusInternalServerError
 func ServeHTTP500(c *gin.Context) {
    c.JSON(http.StatusInternalServerError, nil)
 }

 //501 StatusNotImplemented
 func ServeHTTP501(c *gin.Context) {
    c.JSON(http.StatusNotImplemented, nil)
 }

 //502 StatusBadGateway
 func ServeHTTP502(c *gin.Context) {
    c.JSON(http.StatusBadGateway, nil)
 }

 //503 StatusServiceUnavailable
 func ServeHTTP503(c *gin.Context) {
    c.JSON(http.StatusServiceUnavailable, nil)
 }

 //504 StatusGatewayTimeout
 func ServeHTTP504(c *gin.Context) {
    c.JSON(http.StatusGatewayTimeout, nil)
 }


 //505 StatusHTTPVersionNotSupported
 func ServeHTTP505(c *gin.Context) {
    c.JSON(http.StatusHTTPVersionNotSupported, nil)
 }

 //506 StatusVariantAlsoNegotiates
 func ServeHTTP506(c *gin.Context) {
    c.JSON(http.StatusVariantAlsoNegotiates, nil)
 }

 //507 StatusInsufficientStorage
 func ServeHTTP507(c *gin.Context) {
    c.JSON(http.StatusInsufficientStorage, nil)
 }

 //508 StatusLoopDetected
 func ServeHTTP508(c *gin.Context) {
    c.JSON(http.StatusLoopDetected, nil)
 }

 //510 StatusNotExtended
 func ServeHTTP510(c *gin.Context) {
    c.JSON(http.StatusNotExtended, nil)
 }

 //511 StatusNetworkAuthenticationRequired
 func ServeHTTP511(c *gin.Context) {
    c.JSON(http.StatusNetworkAuthenticationRequired, nil)
 }
