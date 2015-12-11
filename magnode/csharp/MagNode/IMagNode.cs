namespace MagNode {

	public enum ErrNO {
		Success,
		Nil,
	}

	public interface INode {
		ErrNO Init();

		ErrNO DeInit();

		ErrNO Connect(string URL, ulong timeout);

		ErrNO Reconnect(ulong timeout);

		ErrNO Send(byte[] data,uint length,ulong timeout);

		ErrNO Recv(ref byte[] buf, ref uint length,ulong timeout);

		ErrNO Close();

	}
}
