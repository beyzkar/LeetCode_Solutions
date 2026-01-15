package main

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int) //sayı -> index
	//buradaki map sayesinde "hangi sayıları, hangi konumda gördüm?" sorusuna cevap alabiliyoruz

	for i, currentNum := range nums { // range: nums dizisininn içinde başta sona gez
		//i o andaki indeks, currentNum o andaki sayıya eşit oluyor

		x := target - currentNum

		if index, ok := seen[x]; ok { //seen[x]: daha önce x sayısını gördüm mü?
			return []int{index, i} // ok==true x sayısının kaydedilen indexini (index) ve şu andaki indeksi (i) liste halinde döndürüyoruz
		}
		seen[currentNum] = i //ok==false, x map'te yoksa daha sonra kullanılabilmesi için indeksi (i) ve sayıyı (currentNum) map'e kaydediliyor.
	}
	return nil
}
