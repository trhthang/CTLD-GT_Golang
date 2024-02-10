Binary Search (Tìm kiếm nhị phân)

Định nghĩa: 
    * Kỹ thuật tìm kiếm trong 1 mảng đã được sắp xếp
Ý tưởng của thuật toán 
    * Kỹ thuật chia để trị

2 cách cài đặt Binary Search :
    * cài đặt bằng vòng lặp
    * cài đặt bằng đệ quy

Đánh giá thuật toán
    Time Complexity: log(n)
    Space Complexity: 
        * cài đặt bằng vòng lặp: O(1)
        * cài đặt bằng đệ quy: O(log(n))

Khi nào sử dụng Binary Search
    - Khi bài toán có yếu tố sắp xếp
    - Khi thực hiện nhiều truy vấn: Có n truy vấn
        * Tìm kiếm tuần tự: n * TimKiemTuanTu = n*n = O(n^2)
        * Tìm kiếm nhị phân: Sắp xếp + n*TimKiemNhiPhan = nlogn + n*logn = O(nlogn)