package microtrack

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"gomicrotrack/tool"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type ResumenDiario struct {
	XMLName xml.Name   `xml:"http://tempuri.org/ ResumenDiario"`
	Desde   CustomTime `xml:"desde,omitempty"`
	Hasta   CustomTime `xml:"hasta,omitempty"`
	Acceso  string     `xml:"acceso,omitempty"`
}

type ResumenDiarioResponse struct {
	XMLName             xml.Name              `xml:"http://tempuri.org/ ResumenDiarioResponse"`
	ResumenDiarioResult *ArrayOfDatosVehiculo `xml:"ResumenDiarioResult,omitempty"`
}

type ResumenDiarioids struct {
	XMLName xml.Name   `xml:"http://tempuri.org/ ResumenDiarioids"`
	Desde   CustomTime `xml:"desde,omitempty"`
	Hasta   CustomTime `xml:"hasta,omitempty"`
	Acceso  string     `xml:"acceso,omitempty"`
}

type ResumenDiarioidsResponse struct {
	XMLName                xml.Name                `xml:"http://tempuri.org/ ResumenDiarioidsResponse"`
	ResumenDiarioidsResult *ArrayOfDatosVehiculoID `xml:"ResumenDiarioidsResult,omitempty"`
}

type ResumenMensual struct {
	XMLName xml.Name   `xml:"http://tempuri.org/ ResumenMensual"`
	Desde   CustomTime `xml:"desde,omitempty"`
	Hasta   CustomTime `xml:"hasta,omitempty"`
	Acceso  string     `xml:"acceso,omitempty"`
}

type ResumenMensualResponse struct {
	XMLName              xml.Name                     // `xml:"http://tempuri.org/ ResumenMensualResponse"` comento para evitar error de conflicto de tag
	ResumenMensualResult *ArrayOfDatosIdentificadores `xml:"ResumenMensualResult,omitempty"`
}

type UltimaPosicionLlave struct {
	XMLName       xml.Name `xml:"http://tempuri.org/ UltimaPosicionLlave"`
	Identificador string   `xml:"Identificador,omitempty"`
	Acceso        string   `xml:"acceso,omitempty"`
}

type UltimaPosicionLlaveResponse struct {
	XMLName                   xml.Name            `xml:"http://tempuri.org/ UltimaPosicionLlaveResponse"`
	UltimaPosicionLlaveResult *DatosIdentificador `xml:"UltimaPosicionLlaveResult,omitempty"`
}

type ConsultarPosiciones struct {
	XMLName xml.Name   `xml:"http://tempuri.org/ ConsultarPosiciones"`
	Patente string     `xml:"patente,omitempty"`
	Desde   CustomTime `xml:"desde,omitempty"`
	Hasta   CustomTime `xml:"hasta,omitempty"`
	Acceso  string     `xml:"acceso,omitempty"`
}

type ConsultarPosicionesResponse struct {
	XMLName                   xml.Name             `xml:"http://tempuri.org/ ConsultarPosicionesResponse"`
	ConsultarPosicionesResult *ArrayOfItemPosicion `xml:"ConsultarPosicionesResult,omitempty"`
}

type DetalleFranja struct {
	XMLName xml.Name   `xml:"http://tempuri.org/ DetalleFranja"`
	Desde   CustomTime `xml:"desde,omitempty"`
	Hasta   CustomTime `xml:"hasta,omitempty"`
	Acceso  string     `xml:"acceso,omitempty"`
}

type DetalleFranjaResponse struct {
	XMLName             xml.Name            `xml:"http://tempuri.org/ DetalleFranjaResponse"`
	DetalleFranjaResult *ArrayOfDatosFranja `xml:"DetalleFranjaResult,omitempty"`
}

type Char int32
type Duration *Duration
type Guid string
type ArrayOfDatosVehiculo struct {
	XMLName       xml.Name         `xml:"http://schemas.datacontract.org/2004/07/ResumenVehicular ArrayOfDatosVehiculo"`
	DatosVehiculo []*DatosVehiculo `xml:"DatosVehiculo,omitempty"`
}

type DatosVehiculo struct {
	XMLName                   xml.Name   `xml:"http://schemas.datacontract.org/2004/07/ResumenVehicular DatosVehiculo"`
	Conductor                 string     `xml:"Conductor,omitempty"`
	Fecha                     CustomTime `xml:"Fecha,omitempty"`
	Identificador             string     `xml:"Identificador,omitempty"`
	Kms                       float64    `xml:"Kms,omitempty"`
	Maxima                    int32      `xml:"Maxima,omitempty"`
	TiempoActivoSegundos      int32      `xml:"TiempoActivoSegundos,omitempty"`
	TiempoDesconexionSegundos int32      `xml:"TiempoDesconexionSegundos,omitempty"`
	TiempoDetenidoSegundos    int32      `xml:"TiempoDetenidoSegundos,omitempty"`
	TiempoRalentiSegundos     int32      `xml:"TiempoRalentiSegundos,omitempty"`
	Vehiculo                  string     `xml:"Vehiculo,omitempty"`
}

type ArrayOfDatosVehiculoID struct {
	XMLName         xml.Name           `xml:"http://schemas.datacontract.org/2004/07/ResumenVehicular ArrayOfDatosVehiculoID"`
	DatosVehiculoID []*DatosVehiculoID `xml:"DatosVehiculoID,omitempty"`
}

type DatosVehiculoID struct {
	XMLName                   xml.Name   `xml:"http://schemas.datacontract.org/2004/07/ResumenVehicular DatosVehiculoID"`
	Conductor                 string     `xml:"Conductor,omitempty"`
	Fecha                     CustomTime `xml:"Fecha,omitempty"`
	Identificador             string     `xml:"Identificador,omitempty"`
	Kms                       float64    `xml:"Kms,omitempty"`
	Maxima                    int32      `xml:"Maxima,omitempty"`
	TiempoActivoSegundos      int32      `xml:"TiempoActivoSegundos,omitempty"`
	TiempoDesconexionSegundos int32      `xml:"TiempoDesconexionSegundos,omitempty"`
	TiempoDetenidoSegundos    int32      `xml:"TiempoDetenidoSegundos,omitempty"`
	TiempoRalentiSegundos     int32      `xml:"TiempoRalentiSegundos,omitempty"`
	Vehiculo                  string     `xml:"Vehiculo,omitempty"`
	VehiculoID                int32      `xml:"VehiculoID,omitempty"`
	VehiculoPatente           string     `xml:"VehiculoPatente,omitempty"`
}

type ArrayOfDatosIdentificadores struct {
	XMLName              xml.Name                //`xml:"http://schemas.datacontract.org/2004/07/ResumenVehicular ArrayOfDatosIdentificadores"`
	DatosIdentificadores []*DatosIdentificadores `xml:"DatosIdentificadores,omitempty"`
}

type DatosIdentificadores struct {
	XMLName       xml.Name   `xml:"http://schemas.datacontract.org/2004/07/ResumenVehicular DatosIdentificadores"`
	AccCent       int32      `xml:"AccCent,omitempty"`
	Accelerations int32      `xml:"Accelerations,omitempty"`
	Braking       int32      `xml:"Braking,omitempty"`
	FechaDesde    CustomTime `xml:"FechaDesde,omitempty"`
	FechaHasta    CustomTime `xml:"FechaHasta,omitempty"`
	Identificador string     `xml:"Identificador,omitempty"`
	InfHigh       int32      `xml:"InfHigh,omitempty"`
	InfLight      int32      `xml:"InfLight,omitempty"`
	InfMedium     int32      `xml:"InfMedium,omitempty"`
	InfTotal      int32      `xml:"InfTotal,omitempty"`
	Kms           float64    `xml:"Kms,omitempty"`
	Score         float64    `xml:"Score,omitempty"`
	ScoreAP       float64    `xml:"ScoreAP,omitempty"`
	ScoreRR       float64    `xml:"ScoreRR,omitempty"`
	ScoreZU       float64    `xml:"ScoreZU,omitempty"`
}

type DatosIdentificador struct {
	XMLName       xml.Name   `xml:"http://schemas.datacontract.org/2004/07/ResumenVehicular DatosIdentificador"`
	Fecha         CustomTime `xml:"Fecha,omitempty"`
	Identificador string     `xml:"Identificador,omitempty"`
	Latitud       float64    `xml:"Latitud,omitempty"`
	Longitud      float64    `xml:"Longitud,omitempty"`
	Pais          string     `xml:"Pais,omitempty"`
	Vehiculo      string     `xml:"Vehiculo,omitempty"`
}

type ArrayOfItemPosicion struct {
	XMLName      xml.Name        `xml:"http://schemas.datacontract.org/2004/07/ResumenVehicular ArrayOfItemPosicion"`
	ItemPosicion []*ItemPosicion `xml:"ItemPosicion,omitempty"`
}

type ItemPosicion struct {
	XMLName    xml.Name `xml:"http://schemas.datacontract.org/2004/07/ResumenVehicular ItemPosicion"`
	FechaHora  string   `xml:"FechaHora,omitempty"`
	Latitud    float64  `xml:"Latitud,omitempty"`
	Llave      string   `xml:"Llave,omitempty"`
	Longitud   float64  `xml:"Longitud,omitempty"`
	Nombre     string   `xml:"Nombre,omitempty"`
	Operador   string   `xml:"Operador,omitempty"`
	Patente    string   `xml:"Patente,omitempty"`
	TipoEvento string   `xml:"TipoEvento,omitempty"`
	Velocidad  int16    `xml:"Velocidad,omitempty"`
}

type ArrayOfDatosFranja struct {
	DatosFranja []*DatosFranja `xml:"DatosFranja,omitempty"`
}

type DatosFranja struct {
	XMLName          xml.Name   `xml:"http://schemas.datacontract.org/2004/07/ResumenVehicular DatosFranja"`
	DuracionSegundos int32      `xml:"DuracionSegundos,omitempty"`
	Fecha            CustomTime `xml:"Fecha,omitempty"`
	Identificador    string     `xml:"Identificador,omitempty"`
	Latitud          int32      `xml:"Latitud,omitempty"`
	Longitud         int32      `xml:"Longitud,omitempty"`
	Metros           int32      `xml:"Metros,omitempty"`
	Operador         string     `xml:"Operador,omitempty"`
	OperadorID       int32      `xml:"OperadorID,omitempty"`
	Sitio            string     `xml:"Sitio,omitempty"`
	TipoEvento       string     `xml:"TipoEvento,omitempty"`
	Vehiculo         string     `xml:"Vehiculo,omitempty"`
	Zona             string     `xml:"Zona,omitempty"`
}

type IResumenVehicular struct {
	client *SOAPClient
}

func NewIResumenVehicular(url string, tls bool, auth *BasicAuth) *IResumenVehicular {
	if url == "" {
		url = ""
	}
	client := NewSOAPClient(url, tls, auth)
	return &IResumenVehicular{
		client: client,
	}
}

func NewIResumenVehicularWithTLSConfig(url string, tlsCfg *tls.Config, auth *BasicAuth) *IResumenVehicular {
	if url == "" {
		url = ""
	}
	client := NewSOAPClientWithTLSConfig(url, tlsCfg, auth)
	return &IResumenVehicular{
		client: client,
	}
}

func (service *IResumenVehicular) AddHeader(header interface{}) {
	service.client.AddHeader(header)
}

// Backwards-compatible function: use AddHeader instead
func (service *IResumenVehicular) SetHeader(header interface{}) {
	service.client.AddHeader(header)
}

func (service *IResumenVehicular) ResumenDiario(request *ResumenDiario) (*ResumenDiarioResponse, error) {
	response := new(ResumenDiarioResponse)
	err := service.client.Call("http://tempuri.org/IResumenVehicular/ResumenDiario", request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (service *IResumenVehicular) ResumenDiarioids(request *ResumenDiarioids) (*ResumenDiarioidsResponse, error) {
	response := new(ResumenDiarioidsResponse)
	err := service.client.Call("http://tempuri.org/IResumenVehicular/ResumenDiarioids", request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (service *IResumenVehicular) ResumenMensual(request *ResumenMensual) (*ResumenMensualResponse, error) {
	response := new(ResumenMensualResponse)
	err := service.client.Call("http://tempuri.org/IResumenVehicular/ResumenMensual", request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (service *IResumenVehicular) UltimaPosicionLlave(request *UltimaPosicionLlave) (*UltimaPosicionLlaveResponse, error) {
	response := new(UltimaPosicionLlaveResponse)
	err := service.client.Call("http://tempuri.org/IResumenVehicular/UltimaPosicionLlave", request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (service *IResumenVehicular) ConsultarPosiciones(request *ConsultarPosiciones) (*ConsultarPosicionesResponse, error) {
	response := new(ConsultarPosicionesResponse)
	err := service.client.Call("http://tempuri.org/IResumenVehicular/ConsultarPosiciones", request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (service *IResumenVehicular) DetalleFranja(request *DetalleFranja) (error, *DetalleFranjaResponse) {
	response := new(DetalleFranjaResponse)
	err := service.client.Call("http://tempuri.org/IResumenVehicular/DetalleFranja", request, response)
	if err != nil {
		return err, nil
	}
	return nil, response
}

var timeout = time.Duration(30 * time.Second)

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, timeout)
}

type SOAPEnvelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Header  *SOAPHeader
	Body    SOAPBody
}

type SOAPHeader struct {
	XMLName xml.Name      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`
	Items   []interface{} `xml:",omitempty"`
}

type SOAPBody struct {
	XMLName xml.Name    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
	Fault   *SOAPFault  `xml:",omitempty"`
	Content interface{} `xml:",omitempty"`
}

type SOAPFault struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`
	Code    string   `xml:"faultcode,omitempty"`
	String  string   `xml:"faultstring,omitempty"`
	Actor   string   `xml:"faultactor,omitempty"`
	Detail  string   `xml:"detail,omitempty"`
}

const (
	// Predefined WSS namespaces to be used in
	WssNsWSSE string = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd"
	WssNsWSU  string = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd"
	WssNsType string = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordText"
)

type WSSSecurityHeader struct {
	XMLName        xml.Name          `xml:"http://schemas.xmlsoap.org/soap/envelope/ wsse:Security"`
	XmlNSWsse      string            `xml:"xmlns:wsse,attr"`
	MustUnderstand string            `xml:"mustUnderstand,attr,omitempty"`
	Token          *WSSUsernameToken `xml:",omitempty"`
}

type WSSUsernameToken struct {
	XMLName   xml.Name     `xml:"wsse:UsernameToken"`
	XmlNSWsu  string       `xml:"xmlns:wsu,attr"`
	XmlNSWsse string       `xml:"xmlns:wsse,attr"`
	Id        string       `xml:"wsu:Id,attr,omitempty"`
	Username  *WSSUsername `xml:",omitempty"`
	Password  *WSSPassword `xml:",omitempty"`
}

type WSSUsername struct {
	XMLName   xml.Name `xml:"wsse:Username"`
	XmlNSWsse string   `xml:"xmlns:wsse,attr"`
	Data      string   `xml:",chardata"`
}

type WSSPassword struct {
	XMLName   xml.Name `xml:"wsse:Password"`
	XmlNSWsse string   `xml:"xmlns:wsse,attr"`
	XmlNSType string   `xml:"Type,attr"`
	Data      string   `xml:",chardata"`
}

type BasicAuth struct {
	Login    string
	Password string
}

type SOAPClient struct {
	url     string
	tlsCfg  *tls.Config
	auth    *BasicAuth
	headers []interface{}
}

// **********
// Accepted solution from http://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
// Author: Icza - http://stackoverflow.com/users/1705598/icza

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func randStringBytesMaskImprSrc(n int) string {
	src := rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

// **********

func NewWSSSecurityHeader(user, pass, mustUnderstand string) *WSSSecurityHeader {
	hdr := &WSSSecurityHeader{XmlNSWsse: WssNsWSSE, MustUnderstand: mustUnderstand}
	hdr.Token = &WSSUsernameToken{XmlNSWsu: WssNsWSU, XmlNSWsse: WssNsWSSE, Id: "UsernameToken-" + randStringBytesMaskImprSrc(9)}
	hdr.Token.Username = &WSSUsername{XmlNSWsse: WssNsWSSE, Data: user}
	hdr.Token.Password = &WSSPassword{XmlNSWsse: WssNsWSSE, XmlNSType: WssNsType, Data: pass}
	return hdr
}

func (b *SOAPBody) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if b.Content == nil {
		return xml.UnmarshalError("Content must be a pointer to a struct")
	}

	var (
		token    xml.Token
		err      error
		consumed bool
	)

Loop:
	for {
		if token, err = d.Token(); err != nil {
			return err
		}
		if token == nil {
			break
		}
		switch se := token.(type) {
		case xml.StartElement:
			if consumed {
				return xml.UnmarshalError("Found multiple elements inside SOAP body; not wrapped-document/literal WS-I compliant")
			} else if se.Name.Space == "http://schemas.xmlsoap.org/soap/envelope/" && se.Name.Local == "Fault" {
				b.Fault = &SOAPFault{}
				b.Content = nil
				err = d.DecodeElement(b.Fault, &se)
				if err != nil {
					return err
				}
				consumed = true
			} else {
				if err = d.DecodeElement(b.Content, &se); err != nil {
					return err
				}
				consumed = true
			}
		case xml.EndElement:
			break Loop
		}
	}
	return nil
}

func (f *SOAPFault) Error() string {
	return f.String
}

func NewSOAPClient(url string, insecureSkipVerify bool, auth *BasicAuth) *SOAPClient {
	tlsCfg := &tls.Config{
		InsecureSkipVerify: insecureSkipVerify,
	}
	return NewSOAPClientWithTLSConfig(url, tlsCfg, auth)
}

func NewSOAPClientWithTLSConfig(url string, tlsCfg *tls.Config, auth *BasicAuth) *SOAPClient {
	return &SOAPClient{
		url:    url,
		tlsCfg: tlsCfg,
		auth:   auth,
	}
}

func (s *SOAPClient) AddHeader(header interface{}) {
	s.headers = append(s.headers, header)
}

func (s *SOAPClient) Call(soapAction string, request, response interface{}) error {
	envelope := SOAPEnvelope{}
	if s.headers != nil && len(s.headers) > 0 {
		soapHeader := &SOAPHeader{Items: make([]interface{}, len(s.headers))}
		copy(soapHeader.Items, s.headers)
		envelope.Header = soapHeader
	}

	envelope.Body.Content = request
	buffer := new(bytes.Buffer)

	encoder := xml.NewEncoder(buffer)
	//encoder.Indent("  ", "    ")
	if err := encoder.Encode(envelope); err != nil {
		return err
	}
	if err := encoder.Flush(); err != nil {
		return err
	}
	// log.Println(buffer.String())

	req, err := http.NewRequest("POST", s.url, buffer)
	if err != nil {
		return err
	}
	if s.auth != nil {
		req.SetBasicAuth(s.auth.Login, s.auth.Password)
	}

	req.Header.Add("Content-Type", "text/xml; charset=\"utf-8\"")
	req.Header.Add("SOAPAction", soapAction)

	req.Header.Set("User-Agent", "gowsdl/0.1")
	req.Close = true

	tr := &http.Transport{
		TLSClientConfig: s.tlsCfg,
		Dial:            dialTimeout,
	}

	client := &http.Client{Transport: tr}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	rawbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if len(rawbody) == 0 {
		log.Println("empty response")
		return nil
	}

	// log.Println(string(rawbody)) // rawbody tiene el output.
	respEnvelope := new(SOAPEnvelope)
	respEnvelope.Body = SOAPBody{Content: response}
	err = xml.Unmarshal(rawbody, respEnvelope)
	if err != nil {
		return err
	}

	fault := respEnvelope.Body.Fault
	if fault != nil {
		return fault
	}

	return nil
}

type CustomTime struct {
	time.Time
}

func (c *CustomTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	const customForm = "2006-01-02T15:04:05"
	var v string
	err := d.DecodeElement(&v, &start)
	tool.CheckError(err, "info")
	parse, err := time.Parse(customForm, v)
	tool.CheckError(err, "info")
	*c = CustomTime{parse}
	return nil
}

// String converst the time to string.
func (c *CustomTime) String() string {
	s := c.Format("2006-01-02 15:04:05")
	return s
}

func generateTimeWindow(from, to string) (fromt, tot time.Time) {
	fromt, err := time.Parse("2006-01-02 15:04:05", from+" 00:00:00")
	tool.CheckError(err, "panic")
	tot, err = time.Parse("2006-01-02 15:04:05", to+" 23:59:59")
	tool.CheckError(err, "panic")
	return
}

// DayDetail returns a detail of the day. Day must be formatted ISO: YYYYY-MM-DD"
func DayDetail(from, to string) []*DatosFranja {
	PASSWORD := "968b354f5cbc8942d60079b4c1c5620b"
	conector := NewIResumenVehicular("http://dal-serviciosexternosgral.azurewebsites.net/ResumenVehiculo.svc", false, nil)
	fromt, tot := generateTimeWindow(from, to)
	var fromc, toc CustomTime
	fromc, toc = CustomTime{fromt}, CustomTime{tot}
	franja := DetalleFranja{Desde: fromc, Hasta: toc, Acceso: PASSWORD}
	err, out := conector.DetalleFranja(&franja)
	tool.CheckError(err, "panic")
	return out.DetalleFranjaResult.DatosFranja
}

// MonthlyBrief date must be formatted ISO: YYYYY-MM-DD"
func MonthlyBrief(from, to string) *ResumenMensualResponse {
	PASSWORD := "968b354f5cbc8942d60079b4c1c5620b"
	conector := NewIResumenVehicular("http://dal-serviciosexternosgral.azurewebsites.net/ResumenVehiculo.svc", false, nil)
	fromt, tot := generateTimeWindow(from, to)
	out, err := conector.ResumenMensual(&ResumenMensual{
		Desde: CustomTime{
			Time: fromt,
		},
		Hasta: CustomTime{
			Time: tot,
		},
		Acceso: PASSWORD,
	})
	tool.CheckError(err, "panic")
	return out
}
