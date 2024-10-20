/// cat info
type Cat {
    id: String @1

    /// 名称
    name: String @2

    /// 昵称
    nick_name: String @3

    /// 性别, male(Tomcat), female(Queen)
    sex: String @4

    /// 出生日期
    brith_day: Timestamp @5

    /// 地理位置
    location: String @20

    /// 地理范围
    area: String @21

    /// 最后一次发现的时间
    last_found_time: Timestamp @25

    /// 族谱信息
//    family_tree: FamilyTree @40

    /// 状态, trap, neuter, return
    status: String @50

    /// cat的物理或情感状态
    state: String @ 55

    /// create time
    create_time: Timestamp @100

    /// update time
    update_time: Timestamp @101

    /// delete time
    delete_time: Timestamp @102
}
