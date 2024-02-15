

Phân loại:

* Dựa trên số child node

 - Khi số con của 1 node không biết trước --> N-ary Tree
 - Khi số con của 1 node maximum là 3 --> Ternary Tree ( Cây tam phân )
 - Khi số con của 1 node maximum là 2 --> Binary Tree (Cây nhị phân)

* Dựa trên tính chất của cây
 - Cây suy biến: mỗi node chỉ quản lý 1 node --> suy biến thành linked list
 - Cây nhị phân hoàn chỉnh : 1 node quản lý đầy đủ 2 con
 - cây nhị phân gần hoàn chỉnh : nếu bỏ hết tất cả node ở độ cao h, ta thu được 1 cây hoàn chỉnh
 - Trường hợp đặc biệt nhất: Cây nhị phân tìm kiếm (Binary Search Tree) 
    +/ Tất cả các node con bên trái đều nhỏ hơn node hiện tại.
    +/ Tất cả các node con bên phải đều lớn hơn node hiện tại
    ==> Tác dụng: có yếu tố sắp xếp, tạo thuận lợi cho bài toán tìm kiếm

 - AVL Tree (Self-balancing BST) : Cây nhị phân tìm kiếm tự cân bằng
    +/ Cây AVL là cây BST mà tại 1 node bất kỳ chiều cao của cây con trái và cây con phải lệch nhau không quá 1 đơn vị    
    ==> Tác dụng: Giúp cho cây luôn có chiều cao thích hợp để xử lý bài toán


