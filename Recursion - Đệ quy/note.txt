Giới thiệu về đệ quy | Định nghĩa đệ quy 
    - Recursion - Đệ quy: Một hàm tự gọi đến chính nó

Đặc điểm của Đệ quy:
Khi nào sử dụng Đệ quy?
    Một bài toán chỉ áp dụng được phương pháp đệ quy nếu có đủ 2 yếu tố sau đây:
        - Bài toán cơ sở / Điều kiện dừng
        - Công thức quy nạp: Phải đưa về bài toán con nhỏ hơn (cuối cùng là bài toán cơ sở)


Cách chạy của đệ quy có thể được biếu diễn bằng cách sử dụng một cây đệ quy (call tree)


Phân loại đệ quy:
    - Đệ quy thực hiện công việc trước, gọi hàm sau
    - Đệ quy thực hiện công việc sau, goi hàm trước


Khử và tối ưu đệ quy

    * Tối ưu đệ quy: 
        - Đệ quy có nhớ: Lưu lại các kết quả đã tính được trước đó
    * Khử đệ quy: 
        - Tìm cách giải bài toán mà không cần sử dụng đến đệ quy


Khi nào sử dụng đệ quy?

    - Khi chúng ta cần implement nhanh một chức năng nào đó để test hoặc kiểm thử một giải pháp
    mà chưa cần quan tâm đến hiệu năng vội

    - Khi tài nguyên khi sử dụng đệ quy gần tương đương với phương pháp không sử dụng đệ quy
    - Khi đệ quy là cách làm duy nhất mà bạn biết :) 


Hai cách cài đặt đệ quy:
    - Cách 1: không trả về kết quả
        + Áp dụng cho hầy hết các bài toán (nếu có sử dụng đệ quy)
        + Những bài toán hay lưu kết quả (DP, ad-hoc, ...)
    - Cách 2: trả về kết quả
        + Áp dụng cho những CTDL có tính đệ quy(Linked list, tree,...)
        + ngắn gọn, sạch sẽ, dễ hiểu

