package houseform

type HouseAddForm struct {
	Title              string   `json:"title" validate:"required|minLen:2|maxLen:20" message:"required: 请输入标题|minLen:标题至少2个字符|maxLen:标题最多20个字符"`
	Price              uint32   `json:"price" validate:"required|min:1|max:1000000" message:"required: 请输入租金|min:租金范围错误|max:租金范围错误"`
	Storey             uint32   `json:"storey" validate:"required|min:1|max:100" message:"required: 请输入楼层|min:楼层范围错误|max:楼层范围错误"`
	Area               uint32   `json:"area" validate:"required|min:1|max:100000" message:"required: 请输入面积|min:面积范围错误|max:楼层范围错误"`
	Desc               string   `json:"desc" validate:"required|minLen:5|maxLen:1000" message:"required: 请填写房源简介|minLen:简介太短啦|maxLen:简介过长"`
	HouseType          string   `json:"house_type" validate:"required" message:"required: 请选择房源户型"`
	Apartment          string   `json:"apartment" validate:"required" message:"required: 请选择房源户型"`
	Address            string   `json:"address" validate:"required|maxLen:100" message:"required: 请选择房源地址|maxLen:地址过长，换个试试吧"`
	Latitude           string   `json:"latitude" validate:"required" message:"required: 未找到该位置的详细信息，换个地址试试~"`
	Longitude          string   `json:"longitude" validate:"required" message:"required: 未找到该位置的详细信息，换个地址试试~"`
	ShortRent          bool     `json:"short_rent"`
	Resident           string   `json:"resident"`
	Images             []string `json:"img" validate:"required" message:"required: 请上传图片"`
	ProvinceCityRegion []string `json:"region" validate:"required" message:"required: 请选择房源位置"`
	FacilityList       []string `json:"facility_list"`
}
