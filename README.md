
# Tubes_DAP

Kaenova Mahendra Auditama<br>
kaenova@student.telkomuniversity.ac.id<br>
Informatics Engineering, Telkom University, Indonesia<br>
2019

  

## [English]

This program is written in Go Language. This program is formed by giving problems where several needs must be met. I was given the theme of Activity Scheduling. Several specifications must be met, such as:

  
```
1. Create a User Account
    1.1. The program can only store one user account.
    1.2. Input username and password.
    1.3. Username cannot be empty.
    1.4. The password must be at least eight characters in the form of a combination of numbers and letters.
    1.5. The list of activity schedules for this account is left blank.
    1.6. Stored account data in the form of a username and password.
 
2. Login user
    2.1. Create a schedule
        2.1.1. Input the name of the activity, place of activity, start time, and end time.
        2.1.2. The name and place of activity cannot be empty
        2.1.3. Time is the date (day, month, year) and hour (hour, minute)
        2.1.4. Start time > current time
        2.1.5. End time > start time
        2.1.6. There can only be a maximum of 3 scheduled activities at concurrent times.
        2.1.7. ID as the primary key, generated automatically with the format [Current month number in 2 digits] [Last 2 digits of the current year] - [serial number in 3 digits], for example, 0818-010 (meaning the 10th schedule, created in August 2018).
        2.1.8. Stored schedule data in the form of ID, activity name, place of activity, start Time, and end Time.
    2.2. Delete a schedule
        2.2.1. Schedule must be registered.
        2.2.2. Can only delete schedules that have not been missed.
        2.2.3. Displays the schedule for a specific date.
    2.4. Displays the number and data of all schedules that have been made in order based on the start time.
    2.5. Displays the number and schedule data on a particular month (for example, December 2018) in order based on the start time.
    2.6. Displays the number, and schedule data passed, namely the schedule whose completion time <current Time, sorted according to the end time.

```
  
To try it out for and build the file "Ver1.0.go" <br>
We also provide files in the executable form in Release. Only to be run in the command line.

  

## [Bahasa Indonesia]

Program ini dibuat di dalam bahasa Go Language. Program ini dibentuk dengan diberikannya masalah yang dimana ada beberapa kebutuhan yang harus dipenuhi. Saya diberikan tema Penjadwalan Kegiatan. Ada beberapa spesifikasi yang harus dipenuhi, seperti:

  
```
1. Membuat Akun User
	1.1. Program hanya bisa menyimpan satu akun user.
	1.2. Input username dan password.
	1.3. Username tidak boleh kosong.
	1.4. Password minimal 8 karakter berupa kombinasi angka dan huruf.
	1.5. Daftar jadwal kegiatan untuk akun tersebut dikosongkan.
	1.6. Disimpan data akun berupa username dan password.

2. Login user
	2.1. Create jadwal
		2.1.1. Input nama kegiatan, tempat kegiatan, waktu mulai, dan waktu selesai.
		2.1.2. Nama dan tempat kegiatan tidak boleh kosong
		2.1.3. Waktu berupa tanggal (day, month, year) dan jam (hour, minute)
		2.1.4. Waktu mulai > waktu saat ini
		2.1.5. Waktu selesai > waktu mulai
		2.1.6. Hanya boleh ada paling banyak 3 jadwal kegiatan di waktu yang beririsan.
		2.1.7. ID sebagai primary key, digenerate secara otomatis dengan format [nomor bulan saat ini dalam 2 digit][2 digit terakhir dari tahun saat ini]-[nomor urut dalam 3 digit], contoh: 0818-010 (artinya jadwal ke-10, dibuat pada bulan Agustus 2018).
		2.1.8.  Disimpan data jadwal berupa ID, nama kegiatan, tempat kegiatan, waktu mulai, dan waktu selesai.
	2.2. Hapus jadwal
		2.2.1. Jadwal harus sudah terdaftar.
		2.2.2. Hanya bisa menghapus jadwal yang belum terlewat.
		2.2.3. Menampilkan jadwal untuk tanggal tertentu.
	2.4. Menampilkan jumlah dan data semua jadwal yang sudah dibuat secara terurut berdasarkan waktu mulai.
	2.5. Menampilkan jumlah dan data jadwal per-bulan tertentu (misal: Desember 2018) secara terurut berdasarkan waktu mulai.
	2.6. Menampilkan jumlah dan data jadwal yang sudah dilalui, yaitu jadwal yang waktu selesainya < waktu saat ini, secara terurut berdasarkan waktu selesai.

```


<img  align= "right" src="https://cdn.discordapp.com/attachments/527433841690804224/791558706508726292/Pre-comp-3.gif"  width="200">
