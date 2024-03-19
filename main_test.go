
package main
import (
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
)


func SetUpRouter() *gin.Engine{
    gin.SetMode(gin.ReleaseMode)
    router := gin.Default()
    return router
}



func TestServeHTTP100(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/100", ServeHTTP100)
    req, _ := http.NewRequest("GET", "/api/100", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusContinue, w.Code)
}

func TestServeHTTP101(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/101", ServeHTTP101)
    req, _ := http.NewRequest("GET", "/api/101", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusSwitchingProtocols, w.Code)
}

func TestServeHTTP102(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/102", ServeHTTP102)
    req, _ := http.NewRequest("GET", "/api/102", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusProcessing, w.Code)
}

func TestServeHTTP103(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/103", ServeHTTP103)
    req, _ := http.NewRequest("GET", "/api/103", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusEarlyHints, w.Code)
}

func TestHomepageHandler(t *testing.T) {
    r := SetUpRouter()
    r.GET("/", HomepageHandler)
    req, _ := http.NewRequest("GET", "/", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusOK, w.Code)
}

func TestServeHTTP200(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/200", ServeHTTP200)
    req, _ := http.NewRequest("GET", "/api/200", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusOK, w.Code)
}

func TestServeHTTP201(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/201", ServeHTTP201)
    req, _ := http.NewRequest("GET", "/api/201", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusCreated, w.Code)
}

func TestServeHTTP202(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/202", ServeHTTP202)
    req, _ := http.NewRequest("GET", "/api/202", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusAccepted, w.Code)
}

func TestServeHTTP203(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/203", ServeHTTP203)
    req, _ := http.NewRequest("GET", "/api/203", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusNonAuthoritativeInfo, w.Code)
}

func TestServeHTTP204(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/204", ServeHTTP204)
    req, _ := http.NewRequest("GET", "/api/204", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestServeHTTP205(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/205", ServeHTTP205)
    req, _ := http.NewRequest("GET", "/api/205", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusResetContent, w.Code)
}

func TestServeHTTP206(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/206", ServeHTTP206)
    req, _ := http.NewRequest("GET", "/api/206", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusPartialContent, w.Code)
}

func TestServeHTTP207(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/207", ServeHTTP207)
    req, _ := http.NewRequest("GET", "/api/207", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusMultiStatus, w.Code)
}

func TestServeHTTP208(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/208", ServeHTTP208)
    req, _ := http.NewRequest("GET", "/api/208", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusAlreadyReported, w.Code)
}


func TestServeHTTP226(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/226", ServeHTTP226)
    req, _ := http.NewRequest("GET", "/api/226", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusIMUsed, w.Code)
}


func TestServeHTTP300(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/300", ServeHTTP300)
    req, _ := http.NewRequest("GET", "/api/300", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusMultipleChoices, w.Code)
}

func TestServeHTTP301(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/301", ServeHTTP301)
    req, _ := http.NewRequest("GET", "/api/301", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusMovedPermanently, w.Code)
}

func TestServeHTTP302(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/302", ServeHTTP302)
    req, _ := http.NewRequest("GET", "/api/302", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusFound, w.Code)
}

func TestServeHTTP303(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/303", ServeHTTP303)
    req, _ := http.NewRequest("GET", "/api/303", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusSeeOther, w.Code)
}


func TestServeHTTP304(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/304", ServeHTTP304)
    req, _ := http.NewRequest("GET", "/api/304", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusNotModified, w.Code)
}


func TestServeHTTP305(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/305", ServeHTTP305)
    req, _ := http.NewRequest("GET", "/api/305", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusUseProxy, w.Code)
}


func TestServeHTTP307(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/307", ServeHTTP307)
    req, _ := http.NewRequest("GET", "/api/307", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusTemporaryRedirect, w.Code)
}

func TestServeHTTP308(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/308", ServeHTTP308)
    req, _ := http.NewRequest("GET", "/api/308", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusPermanentRedirect, w.Code)
}

func TestServeHTTP400(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/400", ServeHTTP400)
    req, _ := http.NewRequest("GET", "/api/400", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusBadRequest, w.Code)
}


func TestServeHTTP401(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/401", ServeHTTP401)
    req, _ := http.NewRequest("GET", "/api/401", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestServeHTTP402(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/402", ServeHTTP402)
    req, _ := http.NewRequest("GET", "/api/402", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusPaymentRequired, w.Code)
}

func TestServeHTTP403(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/403", ServeHTTP403)
    req, _ := http.NewRequest("GET", "/api/403", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusForbidden, w.Code)
}

func TestServeHTTP404(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/404", ServeHTTP404)
    req, _ := http.NewRequest("GET", "/api/404", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestServeHTTP405(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/405", ServeHTTP405)
    req, _ := http.NewRequest("GET", "/api/405", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusMethodNotAllowed, w.Code)
}



func TestServeHTTP406(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/406", ServeHTTP406)
    req, _ := http.NewRequest("GET", "/api/406", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusNotAcceptable, w.Code)
}

func TestServeHTTP407(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/407", ServeHTTP407)
    req, _ := http.NewRequest("GET", "/api/407", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusProxyAuthRequired, w.Code)
}

func TestServeHTTP408(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/408", ServeHTTP408)
    req, _ := http.NewRequest("GET", "/api/408", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusRequestTimeout, w.Code)
}

func TestServeHTTP409(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/409", ServeHTTP409)
    req, _ := http.NewRequest("GET", "/api/409", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusConflict, w.Code)
}

func TestServeHTTP410(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/410", ServeHTTP410)
    req, _ := http.NewRequest("GET", "/api/410", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusGone, w.Code)
}

func TestServeHTTP411(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/411", ServeHTTP411)
    req, _ := http.NewRequest("GET", "/api/411", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusLengthRequired, w.Code)
}

func TestServeHTTP412(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/412", ServeHTTP412)
    req, _ := http.NewRequest("GET", "/api/412", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusPreconditionFailed, w.Code)
}


func TestServeHTTP413(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/413", ServeHTTP413)
    req, _ := http.NewRequest("GET", "/api/413", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusRequestEntityTooLarge, w.Code)
}

func TestServeHTTP414(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/414", ServeHTTP414)
    req, _ := http.NewRequest("GET", "/api/414", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusRequestURITooLong, w.Code)
}


func TestServeHTTP415(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/415", ServeHTTP415)
    req, _ := http.NewRequest("GET", "/api/415", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusUnsupportedMediaType, w.Code)
}


func TestServeHTTP416(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/416", ServeHTTP416)
    req, _ := http.NewRequest("GET", "/api/416", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusRequestedRangeNotSatisfiable, w.Code)
}

func TestServeHTTP417(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/417", ServeHTTP417)
    req, _ := http.NewRequest("GET", "/api/417", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusExpectationFailed, w.Code)
}

func TestServeHTTP418(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/418", ServeHTTP418)
    req, _ := http.NewRequest("GET", "/api/418", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusTeapot, w.Code)
}

func TestServeHTTP421(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/421", ServeHTTP421)
    req, _ := http.NewRequest("GET", "/api/421", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusMisdirectedRequest, w.Code)
}


func TestServeHTTP422(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/422", ServeHTTP422)
    req, _ := http.NewRequest("GET", "/api/422", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestServeHTTP423(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/423", ServeHTTP423)
    req, _ := http.NewRequest("GET", "/api/423", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusLocked, w.Code)
}

func TestServeHTTP424(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/424", ServeHTTP424)
    req, _ := http.NewRequest("GET", "/api/424", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusFailedDependency, w.Code)
}


func TestServeHTTP425(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/425", ServeHTTP425)
    req, _ := http.NewRequest("GET", "/api/425", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusTooEarly, w.Code)
}

func TestServeHTTP426(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/426", ServeHTTP426)
    req, _ := http.NewRequest("GET", "/api/426", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusUpgradeRequired, w.Code)
}

func TestServeHTTP428(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/428", ServeHTTP428)
    req, _ := http.NewRequest("GET", "/api/428", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusPreconditionRequired, w.Code)
}

func TestServeHTTP429(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/429", ServeHTTP429)
    req, _ := http.NewRequest("GET", "/api/429", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusTooManyRequests, w.Code)
}

func TestServeHTTP431(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/431", ServeHTTP431)
    req, _ := http.NewRequest("GET", "/api/431", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusRequestHeaderFieldsTooLarge, w.Code)
}

func TestServeHTTP451(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/451", ServeHTTP451)
    req, _ := http.NewRequest("GET", "/api/451", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusUnavailableForLegalReasons, w.Code)
}

func TestServeHTTP500(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/500", ServeHTTP500)
    req, _ := http.NewRequest("GET", "/api/500", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusInternalServerError, w.Code)
}


func TestServeHTTP501(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/501", ServeHTTP501)
    req, _ := http.NewRequest("GET", "/api/501", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusNotImplemented, w.Code)
}

func TestServeHTTP502(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/502", ServeHTTP502)
    req, _ := http.NewRequest("GET", "/api/502", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusBadGateway, w.Code)
}

func TestServeHTTP503(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/503", ServeHTTP503)
    req, _ := http.NewRequest("GET", "/api/503", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusServiceUnavailable, w.Code)
}

func TestServeHTTP504(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/504", ServeHTTP504)
    req, _ := http.NewRequest("GET", "/api/504", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusGatewayTimeout, w.Code)
}

func TestServeHTTP505(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/505", ServeHTTP505)
    req, _ := http.NewRequest("GET", "/api/505", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusHTTPVersionNotSupported, w.Code)
}

func TestServeHTTP506(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/506", ServeHTTP506)
    req, _ := http.NewRequest("GET", "/api/506", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusVariantAlsoNegotiates, w.Code)
}

func TestServeHTTP507(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/507", ServeHTTP507)
    req, _ := http.NewRequest("GET", "/api/507", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusInsufficientStorage, w.Code)
}

func TestServeHTTP508(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/508", ServeHTTP508)
    req, _ := http.NewRequest("GET", "/api/508", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusLoopDetected, w.Code)
}

func TestServeHTTP510(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/510", ServeHTTP510)
    req, _ := http.NewRequest("GET", "/api/510", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusNotExtended, w.Code)
}

func TestServeHTTP511(t *testing.T) {
    r := SetUpRouter()

    r.GET("/api/511", ServeHTTP511)
    req, _ := http.NewRequest("GET", "/api/511", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusNetworkAuthenticationRequired, w.Code)
}


