package persongen

import (
	"math/rand"
)

var surnames = []string{"PANTLE", "MASTRODONATO", "TRINGALE", "BORNHOLDT", "PHILIPPSEN", "DROSSART", "GIOCONDO", "ALCIDE", "TISINGER", "WASHER", "HEIDENFELDER", "HELVEY", "BLATTER", "ZOULEK", "MODGLIN", "LIMO", "FINNERN", "CUDAHY", "MOSSI", "AGUINO", "DEHNERT", "CRAFTER", "LAURIANO", "ALBARELLA", "SABOTA", "GHAREEB", "TUTWILER", "GOODSELL", "TRAINHAM", "LEMPICKI", "HANSCOM", "FAUL", "BARCLAY", "SCHOLAR", "SALINS", "EISEMAN", "CRISCI", "BORELAND", "BRUSEWITZ", "KIMBRO", "MAUSBACH", "IANUZZI", "HUDLER", "POLTROCK", "CYBART", "RUND", "ZAHLER", "BALICKI", "CORNELIOUS", "RAFFENSPERGER", "PLASTOW", "LAMPREY", "HOOGEVEEN", "VELAQUEZ", "ASTURIAS", "RUTHE", "LAFORCE", "WILLETS", "MANAOIS", "WECKERLY", "KOLAR", "STADELI", "BOGA", "BOURELLE", "GALIVAN", "VILLEROT", "HEGWER", "SHEMON", "SZADKOWSKI", "MINELLA", "CHELINI", "OCHALEK", "FRUSTER", "MERINO", "EDSTROM", "MURCIA", "HOINS", "BAMMEL", "SOFGE", "LAURELL", "GRASTY", "DEWITT", "SCHNABEL", "SAHAKYAN", "MANDRA", "KOPACEK", "ZEGEL", "SURMA", "FORSTON", "DAVENPORT", "TOSCHI", "RUSHING", "MICHALOWICZ", "GLATT", "FIGLIA", "ZELNO", "SAHL", "SCHARDEIN", "MICHELI", "DAMBECK"}
var firstnames = []string{"Taylor", "Derrick", "Juliana", "Moshe", "Boston", "Will", "Nathan", "Carlo", "Darius", "Micah", "Jordyn", "Ignacio", "Isabel", "Nigel", "Lillie", "Sidney", "Kathryn", "King", "Collin", "Cassius", "Jameson", "Celia", "Gage", "Michaela", "Brandon", "Nikolai", "Salvatore", "Nathanael", "Asher", "James", "Elaina", "Xzavier", "Ethan", "Willie", "Cory", "Edward", "Aiyana", "Kyla", "Rebecca", "Ana", "Sage", "Dallas", "Chase", "Zoe", "Georgia", "Keagan", "Ricardo", "Jennifer", "Jayden", "Landyn", "Koen", "Alena", "Zane", "Samuel", "Porter", "Damion", "Aydin", "Danielle", "Katelyn", "Keira", "Emmett", "Deangelo", "Damarion", "Macey", "Alisa", "Chaya", "Kaila", "Madalynn", "Izayah", "Kailee", "Alyssa", "Jose", "Annabelle", "Nasir", "Keaton", "Callum", "Kenzie", "Gunnar", "Tomas", "Russell", "Evie", "Aniyah", "Dean", "Cristofer", "William", "Franco", "Cynthia", "Neil", "Joanna", "Reagan", "Emerson", "Jaslyn", "Gerardo", "Keon", "Tucker", "Emilio", "Carina", "Peyton", "Isaiah", "Melany"}

func PersonGen(count int) []Person {
	max := 100

	//loop
	var Personlist []Person

	for i := 0; i < count; i++ {
		r := rand.Intn(max - 1)
		var person = Person{
			First: firstnames[r],
			Last:  surnames[r],
		}
		Personlist = append(Personlist, person)
	}
	return Personlist
}
