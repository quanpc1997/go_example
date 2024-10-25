import asyncio
from motor.motor_asyncio import AsyncIOMotorClient
from faker import Faker
from datetime import datetime, timedelta
import random



# Khởi tạo Faker
fake = Faker()

def random_date(start, end):
    """Tạo một ngày ngẫu nhiên giữa start và end."""
    return start + timedelta(days=random.randint(0, (end - start).days))

# Hàm tạo một bản ghi người dùng
def create_user():
    start_date = datetime(1900, 1, 1)
    end_date = datetime(2024, 1, 1)

    # Tạo một ngày ngẫu nhiên
    random_generated_date = random_date(start_date, end_date)
    return {
        "name": fake.name(),
        "username": fake.user_name(),
        "email": fake.email(),
        "address": fake.address(),
        "phone_number": fake.phone_number(),
        "date_of_birth": random_generated_date,
        "job": fake.job(),
        "company": fake.company(),
        "website": fake.url(),
        "bio": fake.text(max_nb_chars=200),
        "city": fake.city(),
        "state": fake.state(),
        "country": fake.country(),
        "zip_code": fake.zipcode(),
        "color": fake.color_name(),
        "language": fake.language_name(),
        "hobby": fake.word(),
        "social_media": {
            "facebook": fake.url(),
            "twitter": fake.url(),
            "linkedin": fake.url(),
        }
    }

async def main():
    # Kết nối tới MongoDB
    client = AsyncIOMotorClient('mongodb://localhost:27017')
    db = client['golang']  # Thay đổi tên cơ sở dữ liệu
    collection = db['users']  # Thay đổi tên collection

    # Tạo và ghi 300k bản ghi người dùng
    users = [create_user() for _ in range(300000)]
    await collection.insert_many(users)

    print("Đã ghi thành công 300,000 bản ghi người dùng vào MongoDB!")

if __name__ == "__main__":
    asyncio.run(main())
