using System.ServiceModel;

namespace GetUserById.BusinessLogic
{
    public class User
    {
        public int Id { get; set; }
        public string? Name { get; set; }
        public string? Email { get; set; }
    }

    [ServiceContract]
    public interface ISoapService
    {
        [OperationContract]
        User GetUserById(int id); 
    }

    public class SoapService : ISoapService
    {
        public User GetUserById(int id)
        {
            // Example admin user
            var users = new List<User>
            {
                new User { Id = 0, Name = "admin", Email = "admin@admin" }
            };

            return users.FirstOrDefault(u => u.Id == id) ?? throw new FaultException($"User with ID {id} not found.");
        }
    }
}