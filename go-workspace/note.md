# **Hướng Dẫn Cài Đặt Linux Centos 7 trên Vmware**
###### POSTED ON 01/22/2021 BY TUNGPHATCOMPUTER 
![HinhAnh1](https://tungphatcomputer.com/wp-content/uploads/2021/01/buoc-18-cai-dat-centos.jpg)
## **Centros là gì ?**
Centos là một hệ điều hành miễn phí được xây dựng và phát triển dựa trên hệ điều hành mã nguồn mở Linux. Centos là chữ viết tắt của “Community Enterprise Operating System”. CentOS ra mắt vào tháng 5/2004 và được phát triển dựa trên bản phân phối của Red Hat Enterprise Linux (RHEL).

Bạn muốn trải nghiệm, tìm hiểu về [hệ điều hành linux](https://tungphatcomputer.com/he-dieu-hanh-linux)  nhưng lại ngại việc cài đặt hệ điều hành linux trực tiếp trên máy tính. Tùng Phát Computer sẽ hướng dẫn các bạn cách cài đặt centos 7 trên máy ảo VMware. 
## **Các bước cài đặt centos 7 trên VMware.**
### **Chuẩn bị:**
* Tải VMware workstation Pro 15.5(<https://www.vmware.com/products/workstation-pro/workstation-pro-evaluation.html>).
* Souce ios centos 7(<www.centos.org/download/>)
* Cấu hình tối thiểu : Ổ cứng : 40 GB , RAM : 4G.
* Cấu hình khuyến cáo : Ổ cứng : 100GB , RAM : 8GB đối với 64bit.

Cài đặt VMware Wokstation

Bước 1: Mở file setup vmware vừa tải về Click chọn Next

![HinhAnh](https://tungphatcomputer.com/wp-content/uploads/2021/01/buoc-1-cai-dat-vmware.jpg)

bước 1 cài đặt vnware

Bước 2: Đánh dấu check vào ô “I accept the terms in the License Agreement”, sau đó tiếp tục bấm Next

![HinhAnh](https://tungphatcomputer.com/wp-content/uploads/2021/01/buoc-2-cai-dat-vmware.jpg)

bước 2 cài đặt vmware

Bước 3: Bước này chọn nơi mà phần mềm được lưu trữ trên máy tính của bạn, nếu không muốn thay đổi thì ta để mặc định, rồi chọn Next.

![HinhAnh](https://tungphatcomputer.com/wp-content/uploads/2021/01/buoc-3-cai-dat-vmware.jpg)

bước 3 cài đặt vnware

Bước 4: Tiếp tục chọn Next

![HinhAnh](https://tungphatcomputer.com/wp-content/uploads/2021/01/buoc-4-cai-dat-vmware.jpg)

bước 4 cài đặt vnware

Bước 5:  Đánh dấu check vào ô “Desktop” để phần mềm tạo một icon ngoài Desktop giúp bạn truy cập chương trình nhanh hơn, sau đó chọn Next.

![HinhAnh](https://tungphatcomputer.com/wp-content/uploads/2021/01/buoc-5-cai-dat-vmware.jpg)

bước 5 cài đặt vnware

Bước 6: Bấm chọn Install để bắt đầu tiến trình cài đặt phần mềm vào máy tính.

![HinhAnh](https://tungphatcomputer.com/wp-content/uploads/2021/01/buoc-5-cai-dat-vmware.jpg)

bước 6 cài đặt vnware

Sau khi cài đặt xong, click chọn Finish để hoàn tất việc cài đặt và bạn đã có thể sử dụng phần mềm.

## **Cài Đặt centos 7**

Sau khi cài đặt VMware xong chúng ta tiến hành cài đặt centos 7 như sau:

Bước 1: Khởi động máy ảo VMware > File > New Virtual Machine.

![](https://tungphatcomputer.com/wp-content/uploads/2021/01/buoc-1-cai-dat-centos.jpg)

Bước 2: Chọn Custom > Next.

![](https://tungphatcomputer.com/wp-content/uploads/2021/01/buoc-1-cai-dat-vmware.jpg)

Bước 3:  Check vào ô I will install the operating system later. > Next

![](https://tungphatcomputer.com/wp-content/uploads/2021/01/buoc-3-cai-dat-centos.jpg)

Bước 4:  Chọn Linux > Version: Centos 7 64-bit (ở đây, máy của mình là 64bit nên mình sẽ chọn là centos 7 64-bit).

![](https://tungphatcomputer.com/wp-content/uploads/2021/01/buoc-4-cai-dat-centos.jpg)

Bước 5:  Đặt tên và chọn nơi lưu file (đây là file đĩa ảo, file này khá nặng nên lưu ở ổ đĩa nào còn trống nhiều dung lượng để tránh gặp lỗi khi cài đặt).

![](https://tungphatcomputer.com/wp-content/uploads/2021/01/buoc-5-cai-dat-centos.jpg)

Bước 6: Chọn tốc độ xử lý của máy ảo.

![](https://tungphatcomputer.com/wp-content/uploads/2021/01/buoc-6-cai-dat-centos.jpg)

Bước 7: Chọn RAM,  nếu ram bạn nhiều thì có thể để lên cao hơn nữa.

![](https://tungphatcomputer.com/wp-content/uploads/2021/01/buoc-7-cai-dat-centos.jpg)

Bước 8: Thiết lập card mạng cho máy ảo ( bạn có thể chọn card nat hoặc bried)

![](https://tungphatcomputer.com/wp-content/uploads/2021/01/buoc-8-cai-dat-centos.jpg)

Bước 9:  Thiết lập chế độ nhập xuất

![](https://tungphatcomputer.com/wp-content/uploads/2021/01/buoc-9-cai-dat-centos.jpg)

Bước 10: tiếp click Next

![](https://tungphatcomputer.com/wp-content/uploads/2021/01/buoc-10-cai-dat-centos.jpg)

Bước 11:  Tạo một ổ đĩa ảo mới. Chọn Create a new virtual disk > Next

![](https://tungphatcomputer.com/wp-content/uploads/2021/01/buoc-11-cai-dat-centos.jpg)

Bước 12: Chọn dung lượng cho ổ đĩa đó, mình để mặc định là 40GB. NHỚ CHỌN Store virtual disk as a single file

![](https://tungphatcomputer.com/wp-content/uploads/2021/01/buoc-12-cai-dat-centos.jpg)

Bước 13: Chọn đường dẫn lưu file ổ đĩa ảo, sau này có thể copy và chuyển qua máy tính khác mã vẫn dùng được máy ảo.

![](https://tungphatcomputer.com/wp-content/uploads/2021/01/buoc-13-cai-dat-centos.jpg)

Bước 14: Chọn Customize Hardware để kiếm tra lại các thiếp lập.

![](https://tungphatcomputer.com/wp-content/uploads/2021/01/buoc-14-cai-dat-centos.jpg)

Bước 15:  New CD/DVD  (IDE) > Use ISO image file và chọn đến đường dẫn file ISO centos 7 đã tải về ở trên. Sau đó ấn Close.

![](https://tungphatcomputer.com/wp-content/uploads/2021/01/buoc-15-cai-dat-centos.jpg)

Bước 16: Chọn Finish

![](https://tungphatcomputer.com/wp-content/uploads/2021/01/buoc-16-cai-dat-centos.jpg)

Bước 17: Khởi động centos bằng Power on this virtual machine

![](https://tungphatcomputer.com/wp-content/uploads/2021/01/buoc-17-cai-dat-centos.jpg)

Bước 18: Chọn Install centos 7

![](https://tungphatcomputer.com/wp-content/uploads/2021/01/buoc-18-cai-dat-centos.jpg)

Bước 19: Thiết lập ngôn ngữ > Continue (english khuyến cáo).

![](https://tungphatcomputer.com/wp-content/uploads/2021/01/buoc-19-cai-dat-centos.jpg)

Bước 20: Thiết lập ngày, giờ. DATE & TIME > Done (set giờ việt nam nhé, cứ trỏ chuột vào bản đồ việt nam)

![](https://tungphatcomputer.com/wp-content/uploads/2021/01/buoc-20-cai-dat-centos.jpg)

Bước 21: Ở mục SOFTWATE SELECTION, chọn giao diện đồ hoạ (GUI) để dễ dàng thao tác:

![](https://tungphatcomputer.com/wp-content/uploads/2021/01/buoc-21-cai-dat-centos.jpg)

Bước 22: Mục INSTALLATION DESTINATION, chọn ổ đĩa ảo 40GB mình đã tạo lúc nãy > Done

![](https://tungphatcomputer.com/wp-content/uploads/2021/01/buoc-22-cai-dat-centos.jpg)

Bước 23: Mục NETWORK & HOST NAME, chọn ON > Done

![](https://tungphatcomputer.com/wp-content/uploads/2021/01/buoc-23-cai-dat-centos.jpg)

Bước 24: Chọn Begin Installation

![](https://tungphatcomputer.com/wp-content/uploads/2021/01/buoc-24-cai-dat-centos.jpg)

Bước 25: Thiết lập tài khoản ROOT, tài khoản này rất quan trọng nên cần phải ghi nhớ mật khẩu

![](https://tungphatcomputer.com/wp-content/uploads/2021/01/buoc-25-cai-dat-centos.jpg)

Bước 26:  Chờ máy tự động cài đặt, bạn chờ khoảng 3 đến 5 phút tuỳ vào cấu hình của máy tính. Click Reboot

![](https://tungphatcomputer.com/wp-content/uploads/2021/01/buoc-26-cai-dat-centos.jpg)

Bước 27: Tick chọn I accept the license agreement

![](https://tungphatcomputer.com/wp-content/uploads/2021/01/buoc-27-cai-dat-centos.jpg)

Bước 28: Tạo tài khoản mới để đăng nhập vào hệ thống.

![](https://tungphatcomputer.com/wp-content/uploads/2021/01/buoc-28-cai-dat-centos.jpg)

Bước 29: Set mật khẩu cho tài khoản của bạn vừa tạo lúc nãy. LƯU Ý 2 MẬT KHẨU ROOT VÀ USER KHÁC NHAU

![](https://tungphatcomputer.com/wp-content/uploads/2021/01/buoc-29-cai-dat-centos.jpg)

Bước 30: Giờ ta đã có thể chạy được Centos trên máy ảo VMware.

![](https://tungphatcomputer.com/wp-content/uploads/2021/01/buoc-30-cai-dat-centos.jpg)

Chúc các bạn cài đặt thành công, ngoài ra để có những trải kiểm tốt hơn trên máy ảo VMware khi cài đặt centos 7 các bạn có thể năng cấp ram của mình hoặc nâng cấp ổ cứng lên. Ngoài ra bạn có thể tham khảo tại tungphatcomputer.com tại đây chúng tôi có cung cấp các [thiết bị máy tính](https://tungphatcomputer.com/may-bo), [ổ cứng SSD](https://tungphatcomputer.com/o-cung-ssd),[HDD](https://tungphatcomputer.com/o-cung-hdd) …, mọi thắc mắc xin liên hệ Tùng Phát Computer.
***
This entry was posted in [Hệ Điều Hành Linux](https://tungphatcomputer.com/he-dieu-hanh-linux). Bookmark the [permalink](https://tungphatcomputer.com/huong-dan-cai-dat-centos-7.html).
***
## **TUNGPHATCOMPUTER**

![](https://secure.gravatar.com/avatar/c286aa83768efb0278f0572d1c8480f9?s=90&d=mm&r=g) | Tùng Phát Computer nơi chia sẻ thủ thuật máy tính PC, Laptop, Mạng,... Cung cấp các giải pháp công nghệ thông tin phục vụ hoạt động các nhân, doanh nghiệp. Call/Zalo 0777 668 568
***
 [Hệ điều hành Linux là gì ? Tìm Hiểu về hệ điều hành Linux](https://tungphatcomputer.com/he-dieu-hanh-linux-la-gi.html) | [Các phiên bản linux dành cho lập trình viên ](https://tungphatcomputer.com/cac-phien-ban-linux.html)
***






