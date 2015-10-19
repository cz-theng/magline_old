namespace MagNode {
	
	public class NodeMgr {
		
		static private NodeMgr instance = null;
		static public NodeMgr Instance
		{
			get
			{
				if(instance == null)
				{
					instance = new NodeMgr();
				}
				
				return instance;
			}
		}
		
		public INode CreateNode() {
			INode node = new Node();
			return node;
		}
		
		public ErrNO DestoryNode(INode node) {
			
			return ErrNO.Success;
		}
	}
}
