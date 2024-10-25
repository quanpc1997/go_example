package service

import (
	"context"
	"example/src/config"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ExportExcel(c *gin.Context) {
	// Kết nối tới MongoDB và collection cụ thể
	collection := config.GetCollection("golang", "users")

	// Thời gian bắt đầu và kết thúc
	startDate := time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)

	TOTAL_RECORD := 300000
	LIMIT_RECORD := 20000
	numer_of_time := (TOTAL_RECORD + LIMIT_RECORD - 1) / LIMIT_RECORD // Tính số lần lặp

	arr := []primitive.M{} // Sử dụng kiểu dữ liệu phù hợp

	for i := 0; i < numer_of_time; i++ {
		skip := i * LIMIT_RECORD
		pipeline := mongo.Pipeline{
			{
				{"$match", bson.D{
					{"date_of_birth", bson.D{
						{"$gte", startDate},
						{"$lte", endDate},
					}},
				}},
			},
			{
				{"$sort", bson.D{{"date_of_birth", -1}}}, // Sắp xếp từ gần nhất đến xa nhất
			},
			{
				{"$skip", skip}, // Bỏ qua số bản ghi đã được xử lý
			},
			{
				{"$limit", LIMIT_RECORD}, // Giới hạn kết quả trả về là 20,000 bản ghi
			},
			{
				{"$group", bson.D{
					{"_id", nil},                         // Không nhóm theo trường nào cả
					{"total_count", bson.D{{"$sum", 1}}}, // Đếm tổng số bản ghi
					{"total_zipcode", bson.D{{"$sum", bson.D{{"$toInt", "$zip_code"}}}}}, // Tính tổng zipcode
					{"records", bson.D{{"$push", "$$ROOT"}}},                             // Đẩy toàn bộ bản ghi vào mảng
				}},
			},
			{
				{"$project", bson.D{
					{"_id", 0},           // Không trả về trường _id
					{"total_count", 1},   // Trả về số lượng tổng
					{"total_zipcode", 1}, // Trả về tổng số zipcode
					{"records", 1},       // Trả về danh sách các bản ghi
				}},
			},
		}

		cursor, err := collection.Aggregate(context.TODO(), pipeline, options.Aggregate())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var results []bson.M
		if err = cursor.All(context.TODO(), &results); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Kiểm tra xem results có phần tử không
		if len(results) > 0 {
			records := results[0]["records"].(primitive.A)
			// Chuyển đổi các phần tử trong records thành kiểu primitive.M và thêm vào arr
			for _, record := range records {
				arr = append(arr, record.(primitive.M))
			}
		}
		if i == (numer_of_time - 1) {
			defer cursor.Close(context.TODO())
		}

	}

	// Thực hiện truy vấn tổng hợp

	count := len(arr)

	// In ra số lượng phần tử trong records
	fmt.Println("Số lượng bản ghi trong results:", count)

	// Trả về kết quả dạng JSON
	c.JSON(http.StatusOK, count) // Trả về arr thay vì results
}
