// TODO 重新设计，方便构建和查询
type FamilyTree {
    cat_id: String @1

    /// 名称
    name: String @2

    /// 昵称
    nick_name: String @3

    /// 性别, male(Tomcat), female(Queen)
    sex: String @4

    /// 祖先
    ancestor: [String] @10

    /// 父母
    parent: [String] @20

    /// 配偶
    spouse: [String] @25

    /// 子女
    child: [String] @30
}
