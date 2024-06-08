package constants

const (
	PROMPT_SUMMARIZE = `Tolong summarize product reviews di bawah ini maksimal 5 kalimat, dan gunakan Bahasa Indonesia:
Output:
{
	"summary": summary
}
	`

	PROMPT_ANALYZE = ` Harap tentukan jumlah komentar di bawah ini yang menggambarkan kondisi baik, buruk, atau tidak memberikan informasi yang cukup di setiap aspek, yaitu: 
	1. packaging atau pengemasan, 
	2. delivery atau pengiriman, 
	3. respon penjual atau respon admin, 
	4. product condition atau kondisi produk. 
	Pastikan bahwa total jumlah komentar tidak melebihi 100, tidak ada yang tumpang tindih, dan setiap komentar hanya diklasifikasikan ke satu kategori. Hanya perlu menampilkan jumlah komentar tanpa rincian komentar aslinya:
jumlah = (komentar positif / (komentar positif + komentar negatif)) * 100
jumlah dalam format float atau bilangan desimal.
Output JSON dengan format sebagai berikut:
{
	"packaging": float(jumlah),
	"delivery": float(jumlah),
	"admin_response": float(jumlah),
	"product_condition": float(jumlah)
}
`
)
