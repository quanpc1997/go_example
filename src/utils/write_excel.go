package utils

import (
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// WriteToExcel nhận vào một mảng các bản ghi và ghi vào file Excel
func WriteToExcel(arr []primitive.M, fileName string) error {
	// Tạo file Excel mới
	f := excelize.NewFile()
	// Tạo một sheet mới và kiểm tra lỗi
	index, err := f.NewSheet("Sheet1")
	if err != nil {
		return fmt.Errorf("lỗi khi tạo sheet: %w", err)
	}

	// Ghi tiêu đề cột
	headers := []string{"Tên", "Tên người dùng", "Email", "Địa chỉ", "Số điện thoại", "Ngày sinh", "Công việc", "Công ty", "Website", "Giới thiệu", "Thành phố", "Tiểu bang", "Quốc gia", "Zip Code", "Màu sắc", "Ngôn ngữ", "Sở thích"}
	for col, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(col+1, 7) // Bắt đầu từ dòng 1
		f.SetCellValue("Sheet1", cell, header)
	}

	// Ghi dữ liệu vào các ô bắt đầu từ dòng 7
	for rowIndex, record := range arr {
		row := rowIndex + 8 // Thay đổi chỉ số dòng bắt đầu từ 7
		cellName, _ := excelize.CoordinatesToCellName(1, row)
		cellUsername, _ := excelize.CoordinatesToCellName(2, row)
		cellEmail, _ := excelize.CoordinatesToCellName(3, row)
		cellAddress, _ := excelize.CoordinatesToCellName(4, row)
		cellPhone, _ := excelize.CoordinatesToCellName(5, row)
		cellDOB, _ := excelize.CoordinatesToCellName(6, row)
		cellJob, _ := excelize.CoordinatesToCellName(7, row)
		cellCompany, _ := excelize.CoordinatesToCellName(8, row)
		cellWebsite, _ := excelize.CoordinatesToCellName(9, row)
		cellBio, _ := excelize.CoordinatesToCellName(10, row)
		cellCity, _ := excelize.CoordinatesToCellName(11, row)
		cellState, _ := excelize.CoordinatesToCellName(12, row)
		cellCountry, _ := excelize.CoordinatesToCellName(13, row)
		cellZip, _ := excelize.CoordinatesToCellName(14, row)
		cellColor, _ := excelize.CoordinatesToCellName(15, row)
		cellLanguage, _ := excelize.CoordinatesToCellName(16, row)
		cellHobby, _ := excelize.CoordinatesToCellName(17, row)

		f.SetCellValue("Sheet1", cellName, record["name"])
		f.SetCellValue("Sheet1", cellUsername, record["username"])
		f.SetCellValue("Sheet1", cellEmail, record["email"])
		f.SetCellValue("Sheet1", cellAddress, record["address"])
		f.SetCellValue("Sheet1", cellPhone, record["phone_number"])

		// Chuyển đổi date_of_birth về định dạng mà Excel có thể hiểu
		if dob, ok := record["date_of_birth"].(primitive.DateTime); ok {
			dobTime := dob.Time() // Chuyển đổi primitive.DateTime thành time.Time
			f.SetCellValue("Sheet1", cellDOB, dobTime.Format("2006-01-02"))
		}

		f.SetCellValue("Sheet1", cellJob, record["job"])
		f.SetCellValue("Sheet1", cellCompany, record["company"])
		f.SetCellValue("Sheet1", cellWebsite, record["website"])
		f.SetCellValue("Sheet1", cellBio, record["bio"])
		f.SetCellValue("Sheet1", cellCity, record["city"])
		f.SetCellValue("Sheet1", cellState, record["state"])
		f.SetCellValue("Sheet1", cellCountry, record["country"])
		f.SetCellValue("Sheet1", cellZip, record["zip_code"])
		f.SetCellValue("Sheet1", cellColor, record["color"])
		f.SetCellValue("Sheet1", cellLanguage, record["language"])
		f.SetCellValue("Sheet1", cellHobby, record["hobby"])
	}

	// Đặt sheet đầu tiên là sheet đang hoạt động
	f.SetActiveSheet(index)

	// Lưu file
	if err := f.SaveAs(fileName); err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("File Excel đã được tạo thành công:", fileName)
	return nil
}
